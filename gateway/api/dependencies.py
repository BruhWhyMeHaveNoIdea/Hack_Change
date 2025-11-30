from fastapi import Request, HTTPException
from auth.jwt import decode_jwt

async def require_auth(request: Request):
    token = request.headers.get("Authorization")
    if not token:
        raise HTTPException(401, "Authorization required")

    try:
        payload = decode_jwt(token.replace("Bearer ", ""))
        request.state.student_id = payload["studentId"]
    except:
        raise HTTPException(401, "Invalid token")