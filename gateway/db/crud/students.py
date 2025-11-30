from db.db import engine
from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy import text, select, delete, update
from db.models.models import Students
from db.schemas.schemas import Students as StudentsDB

async def create_student(student: Students):
    async with AsyncSession(engine) as session:
        student_db = StudentsDB(
            first_name=student.first_name,
            last_name=student.last_name,
            email=student.email,
            password=student.password,
            date=student.date
        )
        session.add(student_db)
        await session.commit()


async def read_student(student_id: int):
    async with AsyncSession(engine) as session:
        result = await session.execute(select(StudentsDB).where(StudentsDB.student_id == student_id)
        )
        query = result.scalars().first()
        return query


async def delete_student(student_id: int):
    async with AsyncSession(engine) as session:
        await session.execute(delete(StudentsDB).where(StudentsDB.student_id == student_id))
        await session.commit()


async def get_all_students():
    async with AsyncSession(engine) as session:
        result = await session.execute(select(StudentsDB))
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