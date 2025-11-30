const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || "http://localhost:8080";

export interface LoginRequest {
  email: string;
  password: string;
}

export interface LoginResponse {
  jwtToken: string;
  refreshToken: string;
}

export interface RefreshTokenRequest {
  refreshToken: string;
}

export interface RefreshTokenResponse {
  jwtToken: string;
}

class ApiClient {
  private baseUrl: string;

  constructor(baseUrl: string) {
    this.baseUrl = baseUrl;
  }

  private async request<T>(
    endpoint: string,
    options: RequestInit = {}
  ): Promise<T> {
    const url = `${this.baseUrl}${endpoint}`;
    const token = this.getToken();

    const headers: HeadersInit = {
      "Content-Type": "application/json",
      ...options.headers,
    };

    if (token) {
      headers["Authorization"] = `Bearer ${token}`;
    }

    const response = await fetch(url, {
      ...options,
      headers,
    });

    if (!response.ok) {
      if (response.status === 401 && token) {
        // Попытка обновить токен только если был токен
        const refreshed = await this.refreshToken();
        if (refreshed) {
          // Повтор запроса с новым токеном
          const newToken = this.getToken();
          if (newToken) {
            headers["Authorization"] = `Bearer ${newToken}`;
            const retryResponse = await fetch(url, {
              ...options,
              headers,
            });
            if (!retryResponse.ok) {
              const errorText = await retryResponse.text();
              throw new Error(
                errorText || `API Error: ${retryResponse.statusText}`
              );
            }
            return retryResponse.json();
          }
        }
        // Если refresh не удался, очищаем токены
        this.clearTokens();
        // Выбрасываем событие для обновления состояния в контексте
        if (typeof window !== "undefined") {
          window.dispatchEvent(new Event("auth:logout"));
        }
        throw new Error("Unauthorized");
      }
      const errorText = await response.text();
      try {
        const errorJson = JSON.parse(errorText);
        throw new Error(errorJson.message || errorJson.error || response.statusText);
      } catch {
        throw new Error(errorText || `API Error: ${response.statusText}`);
      }
    }

    return response.json();
  }

  private getToken(): string | null {
    if (typeof window === "undefined") return null;
    return localStorage.getItem("jwtToken");
  }

  private getRefreshToken(): string | null {
    if (typeof window === "undefined") return null;
    return localStorage.getItem("refreshToken");
  }

  private clearTokens(): void {
    if (typeof window === "undefined") return;
    localStorage.removeItem("jwtToken");
    localStorage.removeItem("refreshToken");
  }

  private setTokens(jwtToken: string, refreshToken: string): void {
    if (typeof window === "undefined") return;
    localStorage.setItem("jwtToken", jwtToken);
    localStorage.setItem("refreshToken", refreshToken);
  }

  async login(credentials: LoginRequest): Promise<LoginResponse> {
    const response = await fetch(`${this.baseUrl}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(credentials),
    });

    if (!response.ok) {
      const error = await response.json().catch(() => ({}));
      throw new Error(error.message || "Login failed");
    }

    const data: LoginResponse = await response.json();
    this.setTokens(data.jwtToken, data.refreshToken);
    return data;
  }

  async refreshToken(): Promise<boolean> {
    const refreshToken = this.getRefreshToken();
    if (!refreshToken) {
      return false;
    }

    try {
      const response = await fetch(`${this.baseUrl}/auth/refresh`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ refreshToken }),
      });

      if (!response.ok) {
        this.clearTokens();
        return false;
      }

      const data: RefreshTokenResponse = await response.json();
      const currentRefreshToken = this.getRefreshToken();
      if (currentRefreshToken) {
        this.setTokens(data.jwtToken, currentRefreshToken);
      }
      return true;
    } catch (error) {
      this.clearTokens();
      return false;
    }
  }

  logout(): void {
    this.clearTokens();
  }

  isAuthenticated(): boolean {
    return !!this.getToken();
  }

  // API методы для курсов
  async getCourses(page: number = 0, size: number = 10) {
    return this.request<any[]>(`/api/courses?page=${page}&size=${size}`);
  }

  async getCourseDetails(courseId: number) {
    return this.request<any>(`/api/courses/${courseId}`);
  }
}

export const apiClient = new ApiClient(API_BASE_URL);

