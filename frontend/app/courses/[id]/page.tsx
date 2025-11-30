"use client";

import { useState, useEffect } from "react";
import { ArrowLeft, AlertCircle, Loader2 } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import Link from "next/link";
import { ChevronDown, ChevronRight } from "lucide-react";
import FooterSection from "@/components/sections/footer/default";
import { use } from "react";
import { apiClient, CourseDetails } from "@/lib/api";
import { useRouter } from "next/navigation";

const getTaskIcon = (type: string) => {
  switch (type) {
    case "code":
      return "💻";
    case "text":
      return "📄";
    case "video":
      return "🎥";
    case "quiz":
      return "❓";
    case "assignment":
      return "📁";
    default:
      return "📝";
  }
};

export default function CoursePage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const [expandedModule, setExpandedModule] = useState<number | null>(null);
  const [course, setCourse] = useState<CourseDetails | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);
  const { id } = use(params);
  const courseId = parseInt(id, 10);
  const router = useRouter();

  useEffect(() => {
    const fetchCourse = async () => {
      try {
        setIsLoading(true);
        setError(null);
        const data = await apiClient.getCourseDetails(courseId);
        setCourse(data);
        // Раскрываем первый модуль по умолчанию
        if (data.modules.length > 0) {
          setExpandedModule(data.modules[0].module_id);
        }
      } catch (err) {
        setError(
          err instanceof Error ? err.message : "Ошибка загрузки курса"
        );
      } finally {
        setIsLoading(false);
      }
    };

    if (courseId) {
      fetchCourse();
    }
  }, [courseId]);

  const toggleModule = (id: number) => {
    setExpandedModule(expandedModule === id ? null : id);
  };

  if (isLoading) {
    return (
      <div className="flex min-h-screen flex-col bg-background">
        <header className="border-b bg-[#003069]">
          <div className="container flex h-16 items-center justify-between py-4">
            <div className="flex items-center gap-3">
              <img
                src="/PSB_white.png"
                alt="ПСБ_White"
                className="h-15 w-auto"
              />
              <span className="text-lg font-semibold text-white">
                Образовательная платформа
              </span>
            </div>
          </div>
        </header>
        <main className="container py-8 px-8 min-w-full flex-grow">
          <div className="flex items-center justify-center py-12">
            <Loader2 className="h-8 w-8 animate-spin text-muted-foreground" />
          </div>
        </main>
        <FooterSection />
      </div>
    );
  }

  if (error || !course) {
    return (
      <div className="flex min-h-screen flex-col bg-background">
        <header className="border-b bg-[#003069]">
          <div className="container flex h-16 items-center justify-between py-4">
            <div className="flex items-center gap-3">
              <img
                src="/PSB_white.png"
                alt="ПСБ_White"
                className="h-15 w-auto"
              />
              <span className="text-lg font-semibold text-white">
                Образовательная платформа
              </span>
            </div>
          </div>
        </header>
        <main className="container py-8 px-8 min-w-full flex-grow">
          <Button variant="ghost" asChild className="mb-6 px-0">
            <Link href="/courses">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Назад к курсам
            </Link>
          </Button>
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Ошибка</AlertTitle>
            <AlertDescription>
              {error || "Курс не найден"}
            </AlertDescription>
          </Alert>
        </main>
        <FooterSection />
      </div>
    );
  }

  return (
    <div className="flex min-h-screen flex-col bg-background">
      <header className="border-b bg-[#003069]">
        <div className="container flex h-16 items-center justify-between py-4">
          <div className="flex items-center gap-3">
            <img src="/PSB_white.png" alt="ПСБ_White" className="h-15 w-auto" />
            <span className="text-lg font-semibold text-white">
              Образовательная платформа
            </span>
          </div>
        </div>
      </header>

      <main className="container py-8 px-8 min-w-full flex-grow">
        <Button variant="ghost" asChild className="mb-6 px-0">
          <Link href="/courses">
            <ArrowLeft className="mr-2 h-4 w-4" />
            Назад к курсам
          </Link>
        </Button>

        <h1 className="text-3xl font-bold mb-4">{course.title}</h1>

        <Card className="mb-8">
          <CardContent className="p-4">
            <p className="text-muted-foreground">{course.description}</p>
          </CardContent>
        </Card>

        <div className="grid grid-cols-1 gap-8 lg:grid-cols-3">
          <div className="space-y-6 lg:col-span-2">
            <Card>
              <CardHeader>
                <CardTitle className="text-lg">Ваш прогресс</CardTitle>
              </CardHeader>
              <CardContent>
                <div className="space-y-2">
                  <Progress value={60} className="h-2" />
                  <p className="text-sm text-muted-foreground">Выполнено 3 из 5 заданий</p>
                </div>
              </CardContent>
            </Card>

            <div className="space-y-4">
              {course.modules
                .sort((a, b) => a.order_index - b.order_index)
                .map((module) => (
                  <Card key={module.module_id}>
                    <CardHeader
                      className="pb-2 cursor-pointer"
                      onClick={() => toggleModule(module.module_id)}
                    >
                      <div className="flex items-center justify-between">
                        <div className="flex items-center gap-2 flex-1">
                          {expandedModule === module.module_id ? (
                            <ChevronDown className="h-5 w-5 text-muted-foreground" />
                          ) : (
                            <ChevronRight className="h-5 w-5 text-muted-foreground" />
                          )}
                          <CardTitle className="text-lg">{module.title}</CardTitle>
                        </div>
                        <Badge variant="secondary">
                          {module.tasks.length} заданий
                        </Badge>
                      </div>
                    </CardHeader>
                    {expandedModule === module.module_id && (
                      <CardContent>
                        <div className="space-y-3">
                          {module.tasks.map((task) => (
                            <div
                              key={task.task_id}
                              className="flex items-start gap-3 p-3 rounded-lg border hover:bg-muted/30 transition-colors"
                            >
                              <span className="text-xl mt-0.5">
                                {getTaskIcon(task.task_type)}
                              </span>
                              <div className="flex-1">
                                <p className="font-medium">{task.title}</p>
                                <p className="text-sm text-muted-foreground">
                                  Тип: {task.task_type} • Баллы: {task.points_value}
                                </p>
                              </div>
                              <Button
                                variant="outline"
                                size="sm"
                                asChild
                              >
                                <Link href={`/tasks/${task.task_id}`}>
                                  Открыть
                                </Link>
                              </Button>
                            </div>
                          ))}
                        </div>
                      </CardContent>
                    )}
                  </Card>
                ))}
            </div>
          </div>

          <div className="lg:col-span-1 space-y-6">
            <Card>
              <CardHeader>
                <CardTitle>Что дальше?</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="text-sm text-muted-foreground">
                  Завершите все задания в модуле, чтобы разблокировать следующий.
                </p>
              </CardContent>
            </Card>
          </div>
        </div>
      </main>
      <FooterSection />
    </div>
  );
}