from pydantic import BaseModel
from datetime import datetime


class Students(BaseModel):
    first_name: str
    last_name: str
    email: str
    password: str
    registration_date: datetime
