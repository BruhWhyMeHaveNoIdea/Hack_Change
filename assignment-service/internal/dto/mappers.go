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

func SubmissionToResponse(s *models.Submission) SubmissionResponse {
	var id, assignment, student string
	if s != nil {
		id = s.ID.String()
		assignment = s.AssignmentID.String()
		student = s.StudentID.String()
	}
	return SubmissionResponse{
		ID:           id,
		AssignmentID: assignment,
		StudentID:    student,
		Content:      s.Content,
		SubmittedAt:  s.SubmittedAt,
		Score:        s.Score,
		GradedAt:     s.GradedAt,
		Feedback:     s.Feedback,
	}
}

func MaterialToResponse(m *models.Material) MaterialResponse {
	var id, course, uploader string
	if m != nil {
		id = m.ID.String()
		course = m.CourseID.String()
		uploader = m.UploaderID.String()
	}
	desc := m.Description
	var storage *string
	if m.StoragePath != nil {
		storage = m.StoragePath
	}
	return MaterialResponse{
		ID:          id,
		CourseID:    course,
		Title:       m.Title,
		Description: desc,
		URL:         m.URL,
		StoragePath: storage,
		UploaderID:  uploader,
		CreatedAt:   m.CreatedAt,
	}
}
