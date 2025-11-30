from fastapi import APIRouter
import requests
import os

router = APIRouter()

@router.get("/health")
async def health_check():
    return {"status": "healthy", "service": "gateway"}

@router.get("/health/detailed")
async def detailed_health_check():
    services_status = {
        "gateway": "healthy",
        "content_service": "unknown", 
        "assignment_service": "unknown"
    }
    
    try:
        content_url = os.getenv("CONTENT_SVC_URL", "http://courses-service:9000")
        response = requests.get(f"{content_url}/health", timeout=5)
        services_status["content_service"] = "healthy" if response.status_code == 200 else "unhealthy"
    except:
        services_status["content_service"] = "unreachable"
    
    try:
        assignment_url = os.getenv("ASSIGNMENT_SVC_URL", "http://assignment-service:8080")
        response = requests.get(f"{assignment_url}/health", timeout=5)
        services_status["assignment_service"] = "healthy" if response.status_code == 200 else "unhealthy"
    except:
        services_status["assignment_service"] = "unreachable"
    
    return services_status

@router.get("/config")
async def show_config():
    return {
        "content_service_url": os.getenv("CONTENT_SVC_URL"),
        "assignment_service_url": os.getenv("ASSIGNMENT_SVC_URL"),
        "gateway_port": os.getenv("GATEWAY_PORT", "8000")
    }