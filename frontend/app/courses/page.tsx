// app/courses/page.tsx
"use client";

import { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { ChevronLeft, ChevronRight, AlertCircle, Loader2 } from "lucide-react";
import Link from "next/link";
import FooterSection from "@/components/sections/footer/default";
import { apiClient, CourseListItem } from "@/lib/api";

const getGradient = (index: number) => {
  const colors = [
    "from-blue-500 to-indigo-600",
    "from-emerald-500 to-teal-600",
    "from-rose-500 to-pink-600",
    "from-amber-500 to-orange-600",
    "from-violet-500 to-purple-600",
    "from-cyan-500 to-sky-600",
    "from-lime-500 to-green-600",
    "from-fuchsia-500 to-purple-600",
    "from-sky-500 to-blue-600",
  ];
  return `bg-gradient-to-r ${colors[index % colors.length]}`;
};

const PAGE_SIZE = 6;

export default function CoursesPage() {
  const [courses, setCourses] = useState<CourseListItem[]>([]);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const [currentPage, setCurrentPage] = useState(0);
  const [totalPages, setTotalPages] = useState(1);

  useEffect(() => {
    const fetchCourses = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const data = await apiClient.getCourses(currentPage, PAGE_SIZE);
        setCourses(data);
        // Предполагаем, что если вернулось меньше PAGE_SIZE, то это последняя страница
        setTotalPages(data.length < PAGE_SIZE ? currentPage + 1 : currentPage + 2);
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Ошибка загрузки курсов"
        );
      } finally {
        setIsLoading(false);
      }
    };

    fetchCourses();
  }, [currentPage]);

  const goToPrev = () => {
    if (currentPage > 0) setCurrentPage(currentPage - 1);
  };

  const goToNext = () => {
    if (currentPage < totalPages - 1) setCurrentPage(currentPage + 1);
  };

  return (
    <div className="flex min-h-screen flex-col bg-background">
      <header className="border-b bg-[#003069] relative overflow-hidden">
        <div className="container flex h-16 items-center justify-between py-4 relative z-10">
          <div className="flex items-center gap-3">
            <img src="/PSB_white.png" alt="ПСБ_White" className="h-15 w-auto" />
            <span className="text-lg font-semibold text-white">
              Образовательная платформа
            </span>
          </div>
        </div>
      </header>

      <main className="container py-8 px-8 min-w-full flex-grow">
        <h1 className="text-3xl font-bold mb-8">Доступные курсы</h1>

        {isLoading ? (
          <div className="flex items-center justify-center py-12">
            <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
          </div>
        ) : error ? (
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Ошибка</AlertTitle>
            <AlertDescription>{error}</AlertDescription>
          </Alert>
        ) : courses.length === 0 ? (
          <Alert>
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Курсы не найдены</AlertTitle>
            <AlertDescription>
              На данный момент нет доступных курсов.
            </AlertDescription>
          </Alert>
        ) : (
          <>
            <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
              {courses.map((course, idx) => {
                const globalIndex = currentPage * PAGE_SIZE + idx;
                return (
                  <Card key={course.id} className="flex flex-col h-full py-0">
                    <div
                      className={`${getGradient(globalIndex)} h-32 rounded-t-lg`}
                    />
                    <CardHeader className="p-4 pb-2">
                      <CardTitle className="">{course.title}</CardTitle>
                    </CardHeader>
                    <CardContent className="p-4 pt-0 space-y-2">
                      <p className="text-sm text-muted-foreground line-clamp-2">
                        {course.description}
                      </p>
                      <span className="inline-block px-2 py-1 text-xs font-medium bg-secondary text-secondary-foreground rounded">
                        {course.difficulty}
                      </span>
                    </CardContent>
                    <CardFooter className="mt-auto p-4 pt-2">
                      <Button variant="outline" className="w-full" asChild>
                        <Link href={`/courses/${course.id}`}>
                          Перейти к курсу
                        </Link>
                      </Button>
                    </CardFooter>
                  </Card>
                );
              })}
            </div>

            <div className="flex items-center justify-between mt-10">
              <Button
                variant="outline"
                onClick={goToPrev}
                disabled={currentPage === 0}
              >
                <ChevronLeft className="h-4 w-4 mr-2" />
                Назад
              </Button>
              <div className="text-sm text-muted-foreground">
                Страница {currentPage + 1} из {totalPages}
              </div>
              <Button
                variant="outline"
                onClick={goToNext}
                disabled={currentPage >= totalPages - 1}
              >
                Вперёд
                <ChevronRight className="h-4 w-4 ml-2" />
              </Button>
            </div>
          </>
        )}
      </main>
      <FooterSection />
    </div>
  );
}
