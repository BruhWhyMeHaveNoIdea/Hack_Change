from datetime import datetime, timedelta
import jwt
from config import JWT_SECRET, JWT_ALGO

def create_jwt(student_id: int):
    payload = {
        "studentId": student_id,
        "exp": datetime.utcnow() + timedelta(minutes=30)
    }
    return jwt.encode(payload, JWT_SECRET, algorithm=JWT_ALGO)

def decode_jwt(token: str):
    return jwt.decode(token, JWT_SECRET, algorithms=[JWT_ALGO])