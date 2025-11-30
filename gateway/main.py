from fastapi import FastAPI
from auth.router import router as auth_router
from api.content_proxy import router as content_router
from api.assignment_proxy import router as assignment_router
from health import router as health_router
from fastapi.middleware.cors import CORSMiddleware

app = FastAPI()

app.add_middleware(
    CORSMiddleware,
    allow_origins=["http://localhost:3000"],
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)


app.include_router(auth_router)
app.include_router(content_router)
app.include_router(assignment_router)
app.include_router(health_router)