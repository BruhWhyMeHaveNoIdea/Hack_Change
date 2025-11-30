/**
 * Утилиты для работы с JWT токенами
 */

export interface TokenPayload {
  exp?: number;
  iat?: number;
  [key: string]: unknown;
}

/**
 * Декодирует JWT токен без проверки подписи
 */
export function decodeToken(token: string): TokenPayload | null {
  try {
    const parts = token.split(".");
    if (parts.length !== 3) {
      return null;
    }

    const payload = parts[1];
    const decoded = atob(payload.replace(/-/g, "+").replace(/_/g, "/"));
    return JSON.parse(decoded);
  } catch (error) {
    console.error("Error decoding token:", error);
    return null;
  }
}

/**
 * Проверяет, истек ли токен
 */
export function isTokenExpired(token: string | null): boolean {
  if (!token) return true;

  const payload = decodeToken(token);
  if (!payload || !payload.exp) return true;

  // exp в секундах, Date.now() в миллисекундах
  const expirationTime = payload.exp * 1000;
  const currentTime = Date.now();

  // Добавляем буфер в 5 минут для обновления токена заранее
  return currentTime >= expirationTime - 5 * 60 * 1000;
}

/**
 * Получает токен из localStorage
 */
export function getToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("jwtToken");
}

/**
 * Получает refresh токен из localStorage
 */
export function getRefreshToken(): string | null {
  if (typeof window === "undefined") return null;
  return localStorage.getItem("refreshToken");
}

/**
 * Сохраняет токены в localStorage
 */
export function setTokens(jwtToken: string, refreshToken: string): void {
  if (typeof window === "undefined") return;
  localStorage.setItem("jwtToken", jwtToken);
  localStorage.setItem("refreshToken", refreshToken);
}

/**
 * Очищает токены из localStorage
 */
export function clearTokens(): void {
  if (typeof window === "undefined") return;
  localStorage.removeItem("jwtToken");
  localStorage.removeItem("refreshToken");
}

/**
 * Проверяет, авторизован ли пользователь
 */
export function isAuthenticated(): boolean {
  const token = getToken();
  if (!token) return false;
  return !isTokenExpired(token);
}


