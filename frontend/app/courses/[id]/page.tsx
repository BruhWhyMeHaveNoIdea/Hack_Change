"use client";

import { useState } from "react";
import { ArrowLeft } from "lucide-react";
import { Button } from "@/components/ui/button";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Progress } from "@/components/ui/progress";
import Link from "next/link";
import { ChevronDown, ChevronRight } from "lucide-react";
import FooterSection from "@/components/sections/footer/default";
import { notFound } from "next/navigation";
import { use } from 'react';

type Task = {
  taskId: number;
  title: string;
  taskType: "Video" | "Reading" | "Quiz" | "Assignment";
  pointsValue: number;
};

type Module = {
  moduleId: number;
  title: string;
  tasks: Task[];
};

type Course = {
  courseId: number;
  title: string;
  description: string;
  modules: Module[];
};

const mockCourse =[{
  courseId: 1,
  title: "Основы SQL и реляционных баз данных",
  description:
    "На этом курсе вы освоите язык SQL, научитесь проектировать реляционные базы данных, писать сложные запросы и оптимизировать производительность.",
  modules: [
    {
      moduleId: 1,
      title: "Введение в базы данных",
      tasks: [
        { taskId: 1, title: "Что такое СУБД?", taskType: "Video", pointsValue: 10 },
        { taskId: 2, title: "Реляционная модель данных", taskType: "Reading", pointsValue: 5 },
        { taskId: 3, title: "Первый SELECT-запрос", taskType: "Quiz", pointsValue: 15 },
      ],
    },
    {
      moduleId: 2,
      title: "Сложные запросы",
      tasks: [
        { taskId: 4, title: "JOIN и его типы", taskType: "Video", pointsValue: 20 },
        { taskId: 5, title: "Подзапросы и CTE", taskType: "Reading", pointsValue: 10 },
        { taskId: 6, title: "Домашнее задание: оптимизация JOIN", taskType: "Assignment", pointsValue: 30 },
      ],
    },
  ],
}];

const getTaskIcon = (type: string) => {
  switch (type) {
    case "Video": return "🎥";
    case "Reading": return "📄";
    case "Quiz": return "❓";
    case "Assignment": return "📁";
    default: return "📝";
  }
};

export default function CoursePage({ params }: { params: Promise<{ id: string }> }) {
    
  const [expandedModule, setExpandedModule] = useState<number | null>(1);
    const { id } = use(params);
  const courseId = parseInt(id, 10);

  const course = mockCourse.find(c => c.courseId === courseId);
  if (!course) notFound();

  const toggleModule = (id: number) => {
    setExpandedModule(expandedModule === id ? null : id);
  };

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
              {course.modules.map((module) => (
                <Card key={module.moduleId}>
                    <CardHeader className="pb-2 cursor-pointer" onClick={() => toggleModule(module.moduleId)}>
                    <div className="flex items-center justify-between">
                        <div className="flex items-center gap-2 flex-1">
                        {expandedModule === module.moduleId ? (
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
                  {expandedModule === module.moduleId && (
                    <CardContent>
                      <div className="space-y-3">
                        {module.tasks.map((task) => (
                          <div
                            key={task.taskId}
                            className="flex items-start gap-3 p-3 rounded-lg border hover:bg-muted/30 transition-colors"
                          >
                            <span className="text-xl mt-0.5">{getTaskIcon(task.taskType)}</span>
                            <div className="flex-1">
                              <p className="font-medium">{task.title}</p>
                              <p className="text-sm text-muted-foreground">
                                Тип: {task.taskType} • Баллы: {task.pointsValue}
                              </p>
                            </div>
                            <Button variant="outline" size="sm">
                              Открыть
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