from fastapi import FastAPI
from auth.router import router as auth_router
from api.content_proxy import router as content_router
from api.assignment_proxy import router as assignment_router
from health import router as health_router  # Импортируем новый роутер

app = FastAPI()


app.include_router(auth_router)
app.include_router(content_router)
app.include_router(assignment_router)
app.include_router(health_router)