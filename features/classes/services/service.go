package services

import (
	"errors"
	"log"
	"mime/multipart"
	"musiclab-be/features/classes"
	"musiclab-be/utils/helper"
	"strings"
)

type classUseCase struct {
	qry classes.ClassData
}

// Delete implements classes.ClassService
func (*classUseCase) Delete(mentorID uint, classID uint) error {
	panic("unimplemented")
}

func New(cd classes.ClassData) classes.ClassService {
	return &classUseCase{
		qry: cd,
	}
}

func (cuc *classUseCase) PostClass(fileData multipart.FileHeader, newClass classes.Core) error {
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return errors.New("validate: " + err.Error())
	}
	newClass.Image = url
	err = cuc.qry.PostClass(newClass)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "data not found"
		} else {
			msg = "server problem"
		}
		log.Println("error add query in service: ", err.Error())
		return errors.New(msg)
	}
	return nil
}

func (cuc *classUseCase) GetMentorClass(mentorID uint) ([]classes.Core, error) {
	res, err := cuc.qry.GetMentorClass(mentorID)

	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "not found") {
			msg = "class not found"
		} else {
			msg = "there is a problem with the server"
		}
		return []classes.Core{}, errors.New(msg)
	}

	return res, nil
}

func (cuc *classUseCase) GetMentorClassDetail(classID uint) (classes.Core, error) {
	res, err := cuc.qry.GetMentorClassDetail(classID)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return classes.Core{}, errors.New("data not found")
		} else {
			return classes.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}

func (cuc *classUseCase) Update(mentorID uint, classID uint, fileData multipart.FileHeader, updatedClass classes.Core) (classes.Core, error) {
	url, err := helper.GetUrlImagesFromAWS(fileData)
	if err != nil {
		return classes.Core{}, errors.New("validate: " + err.Error())
	}
	updatedClass.Image = url
	res, err := cuc.qry.Update(uint(mentorID), classID, updatedClass)

	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return classes.Core{}, errors.New("data not found")
		} else {
			return classes.Core{}, errors.New("internal server error")
		}
	}

	return res, nil
}
