import { NextResponse } from "next/server";
import type { NextRequest } from "next/server";

export function middleware(request: NextRequest) {
  const { pathname } = request.nextUrl;

  // Публичные маршруты (пропускаем проверку)
  const publicRoutes = ["/login", "/"];
  const isPublicRoute = publicRoutes.some((route) => pathname === route);

  // Публичные маршруты пропускаем без проверки
  // Проверка авторизации происходит на клиенте через AuthContext
  if (isPublicRoute) {
    return NextResponse.next();
  }

  // Для остальных маршрутов проверка будет на клиенте
  // Middleware не может проверить localStorage, поэтому пропускаем
  return NextResponse.next();
}

export const config = {
  matcher: [
    /*
     * Match all request paths except for the ones starting with:
     * - api (API routes)
     * - _next/static (static files)
     * - _next/image (image optimization files)
     * - favicon.ico (favicon file)
     * - public files (public folder)
     */
    "/((?!api|_next/static|_next/image|favicon.ico|.*\\.(?:svg|png|jpg|jpeg|gif|webp)$).*)",
  ],
};

