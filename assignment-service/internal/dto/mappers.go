package dto

import (
	"fmt"
	"hack_change/internal/models"
	"time"
)

func UserToResponse(u *models.User) UserResponse {
	var id string
	if u != nil {
		id = u.ID
	}
	return UserResponse{
		ID:         id,
		Email:      u.Email,
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Role:       string(u.Role),
		CreatedAt:  u.CreatedAt,
	}
}

func CourseToResponse(c *models.Course) CourseResponse {
	var id, teacher string
	if c != nil {
		id = c.ID.String()
		teacher = c.TeacherID.String()
	}
	return CourseResponse{
		ID:          id,
		Title:       c.Title,
		Description: c.Description,
		TeacherID:   teacher,
		CreatedAt:   c.CreatedAt,
	}
}

func AssignmentToResponse(a *models.Assignment) AssignmentResponse {
	var aid, sid, tid string
	var upload time.Time
	var fp *string
	var status string
	if a != nil {
		aid = fmt.Sprintf("%d", a.ID)
		sid = a.StudentID
		tid = a.TaskID
		upload = a.UploadDate
		fp = a.FilePath
		status = a.Status
	}
	return AssignmentResponse{
		AssignmentID: aid,
		StudentID:    sid,
		TaskID:       tid,
		UploadDate:   upload,
		FilePath:     fp,
		Status:       status,
	}
}
