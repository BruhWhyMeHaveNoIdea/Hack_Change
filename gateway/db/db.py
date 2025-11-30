from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy import create_engine
from config import DB_CONNECTION_URL

engine = create_engine(DB_CONNECTION_URL, echo=True)
Base = declarative_base()


def create_tables():
    Base.metadata.create_all(engine)