"use client";

import { useState } from "react";
import { useParams } from "next/navigation";
import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Alert, AlertDescription, AlertTitle } from "@/components/ui/alert";
import { Badge } from "@/components/ui/badge";
import { Label } from "@/components/ui/label";
import { Input } from "@/components/ui/input";
import {
  AlertCircle,
  Upload,
  CheckCircle,
  Clock,
  Info,
  ArrowLeft,
} from "lucide-react";
import Link from "next/link";
import FooterSection from "@/components/sections/footer/default";

// Типы
interface Task {
  task_id: number;
  title: string;
  task_info: string;
  task_type: string;
  points_value: number;
}

interface Assignment {
  assignment_id: number;
  status: "Черновик" | "На проверке" | "Проверено";
  upload_date: string;
  file_path: string;
}

interface FeedbackTicket {
  description: string;
  skill: string;
  priority: number;
  status: "To Do" | "In Progress" | "Done";
}

// Моковые данные — ключи как строки!
const mockTasks: Record<string, Task> = {
  "101": {
    task_id: 101,
    title: "Напишите SQL-запрос для соединения таблиц",
    task_info:
      "Используя таблицы users и orders, напишите SELECT-запрос с JOIN, который вернёт имя пользователя и сумму его заказов.",
    task_type: "Assignment",
    points_value: 20,
  },
  "102": {
    task_id: 102,
    title: "Реализуйте REST API на FastAPI",
    task_info:
      "Создайте эндпоинт GET /items, возвращающий список элементов из базы данных.",
    task_type: "Assignment",
    points_value: 25,
  },
};

const mockAssignments: Record<string, Assignment | null> = {
  "101": {
    assignment_id: 501,
    status: "На проверке",
    upload_date: "2025-11-28T14:30:00",
    file_path: "/uploads/sql_homework_v2.sql",
  },
  "102": null,
};

const mockFeedback: Record<string, FeedbackTicket[]> = {
  "101": [
    {
      description:
        "Используйте LEFT JOIN вместо INNER JOIN, чтобы включить пользователей без заказов.",
      skill: "SQL JOINs",
      priority: 1,
      status: "To Do",
    },
    {
      description: "Добавьте алиасы для таблиц для улучшения читаемости.",
      skill: "SQL Style",
      priority: 2,
      status: "Done",
    },
  ],
  "102": [],
};

export default function TaskPage() {
  const params = useParams();
  const taskId = params.taskId as string;

  const task = mockTasks[taskId];
  const assignment = mockAssignments[taskId] ?? null;
  const feedback = mockFeedback[taskId] ?? [];

  const [file, setFile] = useState<File | null>(null);
  const [uploadSuccess, setUploadSuccess] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const handleFileChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.files?.[0]) {
      setFile(e.target.files[0]);
      setError(null);
    }
  };

  const handleUpload = () => {
    if (!file) {
      setError("Выберите файл для загрузки");
      return;
    }
    setUploadSuccess(true);
    setError(null);
    setTimeout(() => {
      alert(`Файл "${file.name}" успешно загружен! (мок)`);
    }, 300);
  };

  if (!task) {
    return (
      <div className="flex min-h-screen flex-col bg-background">
        <header className="border-b bg-[#003069] relative overflow-hidden">
          <div className="container flex h-16 items-center justify-between py-4 relative z-10">
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
        <main className="container py-8 px-8 flex-grow">
          <Alert variant="destructive">
            <AlertCircle className="h-4 w-4" />
            <AlertTitle>Задание не найдено</AlertTitle>
            <AlertDescription>
              Задание с ID {taskId} не существует.
            </AlertDescription>
          </Alert>
          <Button variant="link" asChild className="mt-4">
            <Link href="/courses">Вернуться к курсам</Link>
          </Button>
        </main>
        <FooterSection />
      </div>
    );
  }

  const assignmentStatusMap: Record<string, { label: string; color: string }> =
    {
      Черновик: { label: "Черновик", color: "text-gray-500" },
      "На проверке": { label: "На проверке", color: "text-yellow-500" },
      Проверено: { label: "Проверено", color: "text-green-500" },
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

      {/* Основной контент — двухколоночный */}
      <main className="container py-8 px-8 flex-grow min-w-full">
        <div className="mb-8">
          <Button variant="ghost" asChild className="mb-6 px-0">
            <Link href="/courses">
              <ArrowLeft className="mr-2 h-4 w-4" />
              Назад к курсам
            </Link>
          </Button>
          <div className="flex items-center gap-3 mb-2">
            <Badge variant="secondary">{task.task_type}</Badge>
            <span className="text-sm text-muted-foreground">
              {task.points_value} баллов
            </span>
          </div>
          <h1 className="text-3xl font-bold">{task.title}</h1>
        </div>

        <div className="grid grid-cols-1 xl:grid-cols-3 gap-6">
          {/* Основная колонка — описание, загрузка, фидбек */}
          <div className="xl:col-span-2 space-y-6">
            {/* Описание задания */}
            <Card>
              <CardHeader>
                <CardTitle>Описание задания</CardTitle>
              </CardHeader>
              <CardContent>
                <p className="whitespace-pre-line text-foreground">
                  {task.task_info}
                </p>
              </CardContent>
            </Card>

            {/* Загрузка файла */}
            <Card>
              <CardHeader>
                <CardTitle>Ваша работа</CardTitle>
                {assignment && (
                  <CardDescription>
                    Статус:{" "}
                    <span className="font-medium">{assignment.status}</span> •
                    Загружено:{" "}
                    {new Date(assignment.upload_date).toLocaleDateString(
                      "ru-RU",
                    )}
                  </CardDescription>
                )}
              </CardHeader>
              <CardContent className="space-y-4">
                <div className="space-y-2">
                  <Input
                    id="file-upload"
                    type="file"
                    className="sr-only"
                    onChange={handleFileChange}
                    accept=".pdf,.doc,.docx,.sql,.txt,.zip"
                  />
                  <Label htmlFor="file-upload" className="cursor-pointer">
                    <Button
                      variant="outline"
                      size="sm"
                      className="gap-2"
                      type="button"
                      onClick={() => document.getElementById("file-upload")?.click()}
                    >
                      <Upload className="h-4 w-4" />
                      Выбрать файл
                    </Button>
                  </Label>
                  {file && (
                    <p className="text-sm text-muted-foreground">
                      Выбран файл:{" "}
                      <span className="font-medium">{file.name}</span>
                    </p>
                  )}
                </div>

                <Button
                  onClick={handleUpload}
                  disabled={!file}
                  className="w-full sm:w-auto"
                >
                  Загрузить
                </Button>

                {uploadSuccess && (
                  <Alert
                    variant="default"
                    className="border-green-200 bg-green-50"
                  >
                    <CheckCircle className="h-4 w-4 text-green-600" />
                    <AlertTitle>Успех</AlertTitle>
                    <AlertDescription>
                      Ваш файл был отправлен на проверку.
                    </AlertDescription>
                  </Alert>
                )}
                {error && (
                  <Alert variant="destructive">
                    <AlertCircle className="h-4 w-4" />
                    <AlertDescription>{error}</AlertDescription>
                  </Alert>
                )}
              </CardContent>
            </Card>

            {/* Обратная связь */}
            {feedback.length > 0 && (
              <Card>
                <CardHeader>
                  <CardTitle>Обратная связь</CardTitle>
                  <CardDescription>Замечания от преподавателя</CardDescription>
                </CardHeader>
                <CardContent className="space-y-4">
                  {feedback.map((fb, idx) => (
                    <div
                      key={idx}
                      className="border rounded-md p-4 space-y-2 bg-muted/30"
                    >
                      <div className="flex items-start justify-between">
                        <div className="flex items-start gap-3">
                          {fb.status === "Done" ? (
                            <CheckCircle className="h-5 w-5 text-green-500 mt-0.5 shrink-0" />
                          ) : fb.status === "In Progress" ? (
                            <Clock className="h-5 w-5 text-yellow-500 mt-0.5 shrink-0" />
                          ) : (
                            <Info className="h-5 w-5 text-red-500 mt-0.5 shrink-0" />
                          )}
                          <div>
                            <p className="font-medium">{fb.description}</p>
                            <p className="text-sm text-muted-foreground">
                              Навык: {fb.skill}
                            </p>
                          </div>
                        </div>
                        <Badge
                          variant={
                            fb.status === "Done"
                              ? "secondary"
                              : fb.status === "In Progress"
                                ? "outline"
                                : "destructive"
                          }
                        >
                          {fb.status}
                        </Badge>
                      </div>
                    </div>
                  ))}
                </CardContent>
              </Card>
            )}

            {feedback.length === 0 && assignment?.status === "Проверено" && (
              <Alert className="border-green-200 bg-green-50">
                <CheckCircle className="h-4 w-4 text-green-600" />
                <AlertTitle>Отлично!</AlertTitle>
                <AlertDescription>
                  Преподаватель не нашёл замечаний. Задание выполнено успешно.
                </AlertDescription>
              </Alert>
            )}
          </div>

          {/* Боковая панель — справа (показывается на xl+) */}
          <div className="xl:col-span-1">
            <Card className="sticky top-8">
              <CardHeader>
                <CardTitle>Информация по заданию</CardTitle>
              </CardHeader>
              <CardContent className="space-y-4">
                <div>
                  <p className="text-sm text-muted-foreground">Тип задания</p>
                  <p className="font-medium">{task.task_type}</p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Баллы</p>
                  <p className="font-medium">
                    {task.points_value} из {task.points_value}
                  </p>
                </div>
                <div>
                  <p className="text-sm text-muted-foreground">Статус сдачи</p>
                  {assignment ? (
                    <Badge
                      variant={
                        assignment.status === "Проверено"
                          ? "secondary"
                          : assignment.status === "На проверке"
                            ? "outline"
                            : "default"
                      }
                    >
                      {assignment.status}
                    </Badge>
                  ) : (
                    <Badge variant="destructive">Не отправлено</Badge>
                  )}
                </div>

                <div>
                  <p className="text-sm text-muted-foreground mb-2">
                    Связанные навыки
                  </p>
                  <div className="flex flex-wrap gap-2">
                    {feedback.length > 0 ? (
                      [...new Set(feedback.map((fb) => fb.skill))].map(
                        (skill, i) => (
                          <Badge
                            key={i}
                            variant="secondary"
                            className="text-xs"
                          >
                            {skill}
                          </Badge>
                        ),
                      )
                    ) : (
                      <Badge variant="outline" className="text-xs">
                        Навыки будут указаны после проверки
                      </Badge>
                    )}
                  </div>
                </div>

                <div className="pt-2 border-t">
                  <p className="text-sm text-muted-foreground mb-2">
                    Нужна помощь?
                  </p>
                  <Button variant="outline" size="sm" className="w-full">
                    Связаться с наставником
                  </Button>
                </div>
              </CardContent>
            </Card>
          </div>
        </div>
      </main>

      <FooterSection />
    </div>
  );
}
