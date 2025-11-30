// app/courses/page.tsx
"use client";

import { useState } from "react";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { ChevronLeft, ChevronRight } from "lucide-react";
import Link from "next/link";
import FooterSection from "@/components/sections/footer/default";

const mockCourses = [
  {
    courseId: 1,
    courseName: "Основы SQL и реляционных баз данных",
    difficulty: "Начальный",
  },
  {
    courseId: 2,
    courseName: "Алгоритмы и структуры данных",
    difficulty: "Средний",
  },
  {
    courseId: 3,
    courseName: "Микросервисная архитектура",
    difficulty: "Продвинутый",
  },
  {
    courseId: 4,
    courseName: "Введение в DevOps",
    difficulty: "Начальный",
  },
  {
    courseId: 5,
    courseName: "Безопасность веб-приложений",
    difficulty: "Средний",
  },
  {
    courseId: 6,
    courseName: "Распределённые системы",
    difficulty: "Продвинутый",
  },
  {
    courseId: 7,
    courseName: "CI/CD и автоматизация",
    difficulty: "Средний",
  },
  {
    courseId: 8,
    courseName: "Kubernetes в продакшене",
    difficulty: "Продвинутый",
  },
  {
    courseId: 9,
    courseName: "Введение в Python",
    difficulty: "Начальный",
  },
];

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
  const [currentPage, setCurrentPage] = useState(1);
  const totalPages = Math.ceil(mockCourses.length / PAGE_SIZE);

  const paginatedCourses = mockCourses.slice(
    (currentPage - 1) * PAGE_SIZE,
    currentPage * PAGE_SIZE,
  );

  const goToPrev = () => {
    if (currentPage > 1) setCurrentPage(currentPage - 1);
  };

  const goToNext = () => {
    if (currentPage < totalPages) setCurrentPage(currentPage + 1);
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

        <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-6">
          {paginatedCourses.map((course, idx) => {
            const globalIndex = (currentPage - 1) * PAGE_SIZE + idx;
            return (
              <Card key={course.courseId} className="flex flex-col h-full py-0">
                <div
                  className={`${getGradient(globalIndex)} h-32 rounded-t-lg`}
                />
                <CardHeader className="p-4 pb-2">
                  <CardTitle className="">{course.courseName}</CardTitle>
                </CardHeader>
                <CardContent className="p-4 pt-0">
                  <span className="inline-block px-2 py-1 text-xs font-medium bg-secondary text-secondary-foreground rounded">
                    {course.difficulty}
                  </span>
                </CardContent>
                <CardFooter className="mt-auto p-4 pt-2">
                  <Button variant="outline" className="w-full" asChild>
                    <Link href={`/courses/${course.courseId}`}>
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
            disabled={currentPage === 1}
          >
            <ChevronLeft className="h-4 w-4 mr-2" />
            Назад
          </Button>
          <div className="text-sm text-muted-foreground">
            Страница {currentPage} из {totalPages}
          </div>
          <Button
            variant="outline"
            onClick={goToNext}
            disabled={currentPage === totalPages}
          >
            Вперёд
            <ChevronRight className="h-4 w-4 ml-2" />
          </Button>
        </div>
      </main>
      <FooterSection />
    </div>
  );
}
