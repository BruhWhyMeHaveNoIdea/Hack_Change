"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuth } from "@/contexts/AuthContext";

export default function Home() {
  const router = useRouter();
  const { isAuth, isLoading } = useAuth();

  useEffect(() => {
    if (!isLoading) {
      if (isAuth) {
        router.push("/courses");
      } else {
        router.push("/login");
      }
    }
  }, [isAuth, isLoading, router]);

  return (
    <div className="flex min-h-screen items-center justify-center">
      <div className="text-center">
        <p className="text-muted-foreground">Загрузка...</p>
      </div>
    </div>
  );
}
