from fastapi import FastAPI
from auth.router import router as auth_router
from api.content_proxy import router as content_router
from api.assignment_proxy import router as assignment_router

app = FastAPI()

app.include_router(auth_router)
app.include_router(content_router)
app.include_router(assignment_router)