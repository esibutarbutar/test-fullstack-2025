package main

import (
    "context"
    "crypto/sha1"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "github.com/redis/go-redis/v9"
)

type User struct {
    RealName string `json:"realname"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

var ctx = context.Background()

func sha1Hash(s string) string {
    h := sha1.New()
    h.Write([]byte(s))
    return hex.EncodeToString(h.Sum(nil))
}

func main() {
    rdb := redis.NewClient(&redis.Options{
        Addr: "localhost:6379", 
    })

    app := fiber.New()

    app.Post("/login", func(c *fiber.Ctx) error {
        type LoginRequest struct {
            Username string `json:"username"`
            Password string `json:"password"`
        }
        var req LoginRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(400).JSON(fiber.Map{"error": "Invalid request"})
        }

        key := fmt.Sprintf("login_%s", req.Username)
        val, err := rdb.Get(ctx, key).Result()
        if err == redis.Nil {
            return c.Status(401).JSON(fiber.Map{"error": "User not found"})
        } else if err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Redis error"})
        }

        var user User
        if err := json.Unmarshal([]byte(val), &user); err != nil {
            return c.Status(500).JSON(fiber.Map{"error": "Data error"})
        }

        if user.Password != sha1Hash(req.Password) {
            return c.Status(401).JSON(fiber.Map{"error": "Wrong password"})
        }

        return c.JSON(fiber.Map{
            "message":  "Login successful",
            "realname": user.RealName,
            "email":    user.Email,
        })
    })

    log.Fatal(app.Listen(":3000"))
}