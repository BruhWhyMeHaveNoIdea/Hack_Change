from db.db import Base
from sqlalchemy import Column, Integer, BigInteger, String, Date

class Students(Base):
    __tablename__="students"
    student_id=Column(Integer, autoincrement=True, primary_key=True)
    first_name=Column(String(255))
    last_name=Column(String(255))
    email = Column(String(255))
    password = Column(String(255))
    registration_date = Column(Date)