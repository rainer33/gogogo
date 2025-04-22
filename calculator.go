
package main

import (
    "fmt"
)

func main() {
    var a float64
    var b float64
    var op string

    fmt.Println("🔢 Go 계산기 (종료하려면 Ctrl+C)")
    for {
        fmt.Print("\n계산식을 입력하세요 (예: 3 + 5): ")
        _, err := fmt.Scanf("%f %s %f\n", &a, &op, &b)
        if err != nil {
            fmt.Println("⚠️ 입력 형식이 잘못되었습니다. 예: 3 + 5")
            continue
        }

        switch op {
        case "+":
            fmt.Printf("결과: %.2f\n", a+b)
        case "-":
            fmt.Printf("결과: %.2f\n", a-b)
        case "*":
            fmt.Printf("결과: %.2f\n", a*b)
        case "/":
            if b == 0 {
                fmt.Println("🚫 0으로 나눌 수 없습니다!")
            } else {
                fmt.Printf("결과: %.2f\n", a/b)
            }
        default:
            fmt.Println("❓ 지원하지 않는 연산자입니다.")
        }
    }
}

