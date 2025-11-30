from fastapi import APIRouter, HTTPException
from auth.jwt import create_jwt, decode_jwt
from db.crud.students import get_all_students
import bcrypt


router = APIRouter(prefix="/auth")


@router.post("/login")
async def login(data: dict):
    email = data["email"]
    password = data["password"]
    students = get_all_students()

    students_by_email = {student["email"]: student for student in students}

    if email not in students_by_email or bcrypt.checkpw(password.encode('utf-8'), students_by_email[email]["password"].encode('utf-8')):
        return {
            "status": 401,
            "jwtToken": "",
            "refreshToken": ""
        }

    student_id = students_by_email[email]["student_id"]
    return {
        "status": 200,
        "jwtToken": create_jwt(student_id),
        "refreshToken": create_jwt(student_id)
    }


@router.post("/refresh")
async def refresh(data: dict):
    token = data["refreshToken"]
    payload = decode_jwt(token)

    return {"jwtToken": create_jwt(payload["studentId"])}