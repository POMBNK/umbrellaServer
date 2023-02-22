## Task
Build a middleware using [echo framework](https://echo.labstack.com/)

## Endpoints
- Create handler, which sends how many days left before 1st Jan 2025 and response `HTTP 200 OK` status code.
- Build middleware that checks HTTP header `User-Role` presents and contain admin and prints `"red button user detected"` to the console.
- Think about coding standarts and best practices. Try to create a good code-structure.

## Conclusion

- Used [this source](https://github.com/golang-standards/project-layout) as a main project layout
- Echo has method `Start()` to run http server in a handler domain. So, if you want, you can create a custom server from "net/http" using a *http.Server, because `Echo pkg` also use it under hood. 
- Made a custom server example for app, that you can see in `cmd/customapp/`.
- Used a DI approach for low code coherence 
- Added graceful shutdown to both servers