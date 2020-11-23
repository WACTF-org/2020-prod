using app.Middleware;
using app.Models;
using Microsoft.AspNetCore.Mvc;

namespace app.Controllers
{
    public class HomeController : Controller
    {
        public IActionResult Index()
        {
            GetSessionData();
            return View();
        }

        public IActionResult Admin()
        {
            GetSessionData();
            if (ViewData["isadmin"] != null && (bool)ViewData["isadmin"] == true)
            {
                return View();
            }
            else
            {
                return RedirectToAction("AccessDenied");
            }
        }

        public IActionResult SessionInspector()
        {
            GetSessionData();
            return View();
        }

        public IActionResult Login()
        {
            GetSessionData();
            return View();
        }

        public IActionResult Logout()
        {
            var newSession = new SecureSession
            {
                Username = "guest",
                IsAdmin = false,
            };
            SecureSessionMiddleware.SetSession(HttpContext, newSession);
            return RedirectToAction("Index");
        }

        [HttpPost]
        public IActionResult Login([Bind("Username,Password")] Login loginModel)
        {
            GetSessionData();
            if (loginModel.Username == "admin" && loginModel.Password == "f1f5f48f8706d4ff64ef7dc48b9ba01816e9a575ec48262cfb9fbee5e4b39fa8")
            {
                var newSession = new SecureSession
                {
                    Username = "admin",
                    IsAdmin = true,
                };
                SecureSessionMiddleware.SetSession(HttpContext, newSession);
                return RedirectToAction("Index");
            }
            else
            {
                ViewData.Add("error", true);
                return View();
            }
        }

        public IActionResult AccessDenied()
        {
            GetSessionData();
            return View();
        }

        private void GetSessionData()
        {
            object session;
            HttpContext.Items.TryGetValue("SESSION", out session);
            if (session != null)
            {
                ViewData.Add("username", ((SecureSession)session).Username);
                ViewData.Add("isadmin", ((SecureSession)session).IsAdmin);
            }
        }
    }
}
