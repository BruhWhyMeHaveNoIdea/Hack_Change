import httpx

async def forward_request(
    method: str,
    url: str,
    headers: dict = None,
    json=None,
    data=None,
    files=None,
):
    async with httpx.AsyncClient() as client:
        response = await client.request(
            method,
            url,
            headers=headers,
            json=json,
            data=data,
            files=files
        )

    return response