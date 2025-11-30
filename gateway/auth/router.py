from fastapi import APIRouter, HTTPException
from auth.jwt import create_jwt, decode_jwt
from db.crud.students import get_all_students


router = APIRouter(prefix="/auth")


@router.post("/login")
async def login(data: dict):
    email = data["email"]
    password = data["password"]
    students = await get_all_students()

    students_by_email = {student["email"]: student for student in students}

    if email not in students_by_email or students_by_email[email]["password"] != password:
        raise HTTPException(status_code=401, detail="Invalid credentials")

    student_id = students_by_email[email]["student_id"]
    return {
        "jwtToken": create_jwt(student_id),
        "refreshToken": create_jwt(student_id)
    }


@router.post("/refresh")
async def refresh(data: dict):
    token = data["refreshToken"]
    payload = decode_jwt(token)

    return {"jwtToken": create_jwt(payload["studentId"])}