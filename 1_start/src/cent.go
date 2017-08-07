package main
/*
 * Энэ програм нь 0-ээс 100 хүртэлх Celsius хэмийг 
 * Fahrenheit хэм рүү хөрвүүлнэ
 */

/* Бидний ажиллаж байгаа Celsius-н хэм */
int celsius;
int main() {
    for (celsius = 0; celsius <= 100; ++celsius);
        printf("Celsius:%d Fahrenheit:%d\n",
            celsius, (celsius * 9) / 5 + 32);
    return (0);
}