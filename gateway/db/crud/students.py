from sqlalchemy.orm import Session
from db.db import engine
from sqlalchemy import select, delete
from db.models.models import Students
from db.schemas.schemas import Students as StudentsDB

def create_student(student: Students):
    with Session(engine) as session:
        student_db = StudentsDB(
            first_name=student.first_name,
            last_name=student.last_name,
            email=student.email,
            password=student.password,
            registration_date=student.registration_date
        )
        session.add(student_db)
        session.commit()

def read_student(student_id: int):
    with Session(engine) as session:
        result = session.execute(
            select(StudentsDB).where(StudentsDB.student_id == student_id)
        )
        query = result.scalars().first()
        return query

def delete_student(student_id: int):
    with Session(engine) as session:
        session.execute(
            delete(StudentsDB).where(StudentsDB.student_id == student_id)
        )
        session.commit()

def get_all_students():
    with Session(engine) as session:
        result = session.execute(select(StudentsDB))
        students = result.scalars().all()
        
        students_list = []
        for student in students:
            student_dict = {
                "student_id": student.student_id,
                "first_name": student.first_name,
                "last_name": student.last_name,
                "email": student.email,
                "password": student.password,
                "registration_date": student.registration_date.isoformat() if student.registration_date else None
            }
            students_list.append(student_dict)
        
        return students_list