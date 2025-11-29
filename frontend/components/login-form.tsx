"use client";

import { zodResolver } from "@hookform/resolvers/zod";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { Button } from "@/components/ui/button";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";

const formSchema = z.object({
  email: z.string().email({ message: "Некорректный email" }),
  password: z.string().min(1, { message: "Обязательное поле" }),
});

export function LoginForm() {
  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      email: "",
      password: "",
    },
  });

  async function onSubmit(values: z.infer<typeof formSchema>) {
    console.log("Login attempt:", values);
  }

  return (
    <div className="flex min-h-screen min-w-full items-center justify-center bg-background p-4">
      <Card className="w-full max-w-md">
        <CardHeader>
        <div className="flex justify-center mb-6">
            <img 
                src="/PSB.png" 
                alt="ПСБ" 
                className="h-12 w-auto" 
            />
        </div>
        <CardTitle className="text-2xl text-center">Вход в образовательную платформу</CardTitle>
        </CardHeader>
        <CardContent>
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-6">
              <FormField
                control={form.control}
                name="email"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Электронная почта</FormLabel>
                    <FormControl>
                      <Input placeholder="ivan@example.com" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="password"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Пароль</FormLabel>
                    <FormControl>
                      <Input type="password" {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <Button
                type="submit"
                className="w-full mt-6"
                disabled={form.formState.isSubmitting}
              >
                {form.formState.isSubmitting ? "Вход..." : "Войти"}
              </Button>
            </form>
          </Form>
        </CardContent>
      </Card>
    </div>
  );
}