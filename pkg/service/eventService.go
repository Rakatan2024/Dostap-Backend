package service

import (
	"hellowWorldDeploy/pkg/entity"
	"math/rand"
	"mime/multipart"
	"time"
)

type EventInterface interface {
	CreateEvent(event *entity.Event, fileHeaders []*multipart.FileHeader) error
	GetSpecialForUserEvents(int, string) ([]entity.Event, error)
	GetAllEvents() ([]entity.Event, error)
	GetEventsByPage(limit, offset int) ([]entity.Event, error)
}

func (s *Service) CreateEvent(event *entity.Event, fileHeaders []*multipart.FileHeader) error {
	pictureDirectoryLink := s.generateRandomKey(entity.PicturesLinkNameLength)
	event.Link = pictureDirectoryLink
	err := s.repo.CreateEvent(event)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service): %s\n", err.Error())
		return err
	}
	err = s.repo.CreateEventOrganizers(event.ID, event.OrganizerUsernames)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service) during CreateEventOrganizers : %s\n", err.Error())
		return err
	}
	err = s.repo.CreateEventInterests(event.ID, event.InterestIDs)
	if err != nil {
		s.log.Printf("\nError CreateEvent(service) during CreateEventInterests : %s\n", err.Error())
		return err
	}
	for _, fileHeader := range fileHeaders {
		file, err := fileHeader.Open()
		if err != nil {
			s.log.Printf("\nError CreateEvent(service) in opening files : %s\n", err.Error())
			return err
		}
		if err := s.UploadFile(file, pictureDirectoryLink+"/"+s.generateRandomKey(entity.PicturesLinkNameLength)); err != nil {
			s.log.Printf("error during upload file in CreateEvent(Service) %s", err.Error())
			return err
		}
		err = file.Close()
		if err != nil {
			s.log.Printf("\nError CreateEvent(service) in closing files : %s\n", err.Error())
			return err
		}
	}
	//log.Println(event.ID, "----Service")
	return nil
}

func (s *Service) GetSpecialForUserEvents(userId int, city string) ([]entity.Event, error) {
	userInterests, err := s.repo.GetUserInterestIdsArray(userId)
	if err != nil {
		s.log.Printf("\nError GetSpecialForUserEvents(service): %s\n", err.Error())
		return nil, err
	}
	events, err := s.repo.GetEventsByInterests(userInterests, city)
	if err != nil {
		s.log.Printf("\nError GetSpecialForUserEvents(service): %s\n", err.Error())
		return nil, err
	}
	return events, nil
}

func (s *Service) GetAllEvents() ([]entity.Event, error) {
	events, err := s.repo.GetAllEvents()
	for i := range events {
		if links, err := s.bc.ListObjects(entity.BucketName, events[i].Link); err != nil {
			return nil, err
		} else {
			events[i].Links = links
		}
	}
	if err != nil {
		s.log.Printf("\nError GetAllEvents(service): %s\n", err.Error())
		return nil, err
	}
	return events, nil
}

func (s *Service) UploadFile(file multipart.File, link string) error {
	err := s.bc.UploadFile(entity.BucketName, link, file)
	if err != nil {
		s.log.Printf("\nerror during UploadFile(Service): %s\n", err.Error())
	}
	return err
}

func (s *Service) GetEventsByPage(limit, offset int) ([]entity.Event, error) {
	events, err := s.repo.GetEventsByPage(limit, offset)
	for i := range events {
		if links, err := s.bc.ListObjects(entity.BucketName, events[i].Link); err != nil {
			return nil, err
		} else {
			events[i].Links = links
		}
	}
	if err != nil {
		s.log.Printf("\nerror during GetEventsPage(Service): %s\n", err.Error())
		return nil, err
	}
	return events, nil
}

func (s *Service) generateRandomKey(length int) string {
	// Use current time as the seed
	rand.New(rand.NewSource(time.Now().UnixNano()))
	key := make([]byte, length)
	for i := range key {
		key[i] = entity.Charset[rand.Intn(len(entity.Charset))]
	}
	return string(key)
}
