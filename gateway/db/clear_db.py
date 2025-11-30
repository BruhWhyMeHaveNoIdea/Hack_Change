from sqlalchemy.ext.asyncio import AsyncSession
from sqlalchemy.sql import text
from db import engine
import asyncio


async def get_all_tables():
    async with AsyncSession(engine) as session:
        result = await session.execute(
            text("""
                SELECT table_name 
                FROM information_schema.tables 
                WHERE table_schema = 'public'
            """)
        )
        tables = result.fetchall()
        return [table[0] for table in tables]


async def delete_tables():
    async with AsyncSession(engine) as session:
        tables = await get_all_tables()
        for table in tables:
            await session.execute(text(f"DROP TABLE IF EXISTS {table}"))
        await session.commit()


if __name__ == "__main__":
    asyncio.run(delete_tables())