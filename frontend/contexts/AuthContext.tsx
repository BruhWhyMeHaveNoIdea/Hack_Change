"use client";

import React, { createContext, useContext, useState, useEffect } from "react";
import { useRouter, usePathname } from "next/navigation";
import { apiClient } from "@/lib/api";
import { isAuthenticated, isTokenExpired, getToken } from "@/lib/auth";

interface AuthContextType {
  isAuth: boolean;
  isLoading: boolean;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
  checkAuth: () => Promise<boolean>;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [isAuth, setIsAuth] = useState(false);
  const [isLoading, setIsLoading] = useState(true);
  const router = useRouter();
  const pathname = usePathname();

  const checkAuth = async (): Promise<boolean> => {
    try {
      const token = getToken();
      if (!token) {
        setIsAuth(false);
        return false;
      }

      // Проверяем, не истек ли токен
      if (isTokenExpired(token)) {
        // Пытаемся обновить токен
        const refreshed = await apiClient.refreshToken();
        if (refreshed) {
          setIsAuth(true);
          return true;
        } else {
          setIsAuth(false);
          apiClient.logout();
          return false;
        }
      }

      setIsAuth(true);
      return true;
    } catch (error) {
      console.error("Auth check error:", error);
      setIsAuth(false);
      return false;
    }
  };

  const login = async (email: string, password: string): Promise<void> => {
    try {
      await apiClient.login({ email, password });
      setIsAuth(true);
      router.push("/courses");
    } catch (error) {
      console.error("Login error:", error);
      throw error;
    }
  };

  const logout = () => {
    apiClient.logout();
    setIsAuth(false);
    router.push("/login");
  };

  useEffect(() => {
    let mounted = true;

    const initAuth = async () => {
      setIsLoading(true);
      const authenticated = await checkAuth();
      
      if (!mounted) return;
      
      setIsLoading(false);

      // Редирект логика только при первой загрузке
      const isLoginPage = pathname === "/login";
      const isHomePage = pathname === "/";

      if (isHomePage) {
        // Главная страница - редирект в зависимости от авторизации
        if (authenticated) {
          router.replace("/courses");
        } else {
          router.replace("/login");
        }
      } else if (!authenticated && !isLoginPage) {
        // Не авторизован и не на странице логина - редирект на логин
        router.replace("/login");
      } else if (authenticated && isLoginPage) {
        // Авторизован и на странице логина - редирект на курсы
        router.replace("/courses");
      }
    };

    initAuth();

    return () => {
      mounted = false;
    };
  }, []); // Только при монтировании

  // Периодическая проверка токена
  useEffect(() => {
    if (!isAuth) return;

    const interval = setInterval(async () => {
      const token = getToken();
      if (token && isTokenExpired(token)) {
        const refreshed = await apiClient.refreshToken();
        if (!refreshed) {
          setIsAuth(false);
          router.push("/login");
        }
      }
    }, 5 * 60 * 1000); // Проверка каждые 5 минут

    // Обработка события выхода из системы
    const handleLogout = () => {
      setIsAuth(false);
      router.push("/login");
    };

    window.addEventListener("auth:logout", handleLogout);

    return () => {
      clearInterval(interval);
      window.removeEventListener("auth:logout", handleLogout);
    };
  }, [isAuth, router]);

  return (
    <AuthContext.Provider value={{ isAuth, isLoading, login, logout, checkAuth }}>
      {children}
    </AuthContext.Provider>
  );
}

export function useAuth() {
  const context = useContext(AuthContext);
  if (context === undefined) {
    throw new Error("useAuth must be used within an AuthProvider");
  }
  return context;
}
