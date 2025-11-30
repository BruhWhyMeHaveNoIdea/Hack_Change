import dotenv
import os

dotenv.load_dotenv(".env")

CONTENT_SVC_URL = os.getenv("CONTENT_SVC_URL")
ASSIGNMENT_SVC_URL = os.getenv("ASSIGNMENT_SVC_URL")
JWT_SECRET_KEY = os.getenv("JWT_SECRET_KEY")
JWT_ALGO = os.getenv("JWT_ALG")

DB_CONNECTION_URL = os.getenv("SQLALCH_DATABASE_URL")