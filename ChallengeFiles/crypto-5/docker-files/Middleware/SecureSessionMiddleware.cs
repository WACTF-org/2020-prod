using Microsoft.AspNetCore.Http;
using System.Threading.Tasks;
using System.Security.Cryptography;
using System.Text;
using System;
using app.Models;
using System.Text.Json;

namespace app.Middleware
{
    public class SecureSessionMiddleware
    {
        private readonly RequestDelegate _next;
        private static readonly RNGCryptoServiceProvider _rng = new RNGCryptoServiceProvider();
        private static readonly byte[] _key = Encoding.UTF8.GetBytes("pTp3RxxaFhAsutVkknaZifttuknduito");

        public SecureSessionMiddleware(RequestDelegate next)
        {
            _next = next;
        }

        public async Task InvokeAsync(HttpContext context)
        {
            string encryptedCookie = context.Request.Cookies["SecureSession"];
            if (string.IsNullOrWhiteSpace(encryptedCookie))
            {
                var session = new SecureSession
                {
                    Username = "guest",
                    IsAdmin = false,
                };
                SetSession(context, session);
            }
            else
            {
                var session = DeserializeSession(DecryptCookie(encryptedCookie));
                context.Items.Add("SESSION", session);
            }

            await _next(context);
        }

        public static void SetSession(HttpContext context, SecureSession session)
        {
            var cookieValue = EncryptSession(SerializeSession(session));
            context.Response.Cookies.Delete("SecureSession");
            context.Response.Cookies.Append("SecureSession", cookieValue);
        }

        private static byte[] DecryptCookie(string cookie)
        {
            byte[] decrypted = null;
            var cookieSplit = cookie.Split(new char[] {'|'});
            var iv = Convert.FromBase64String(cookieSplit[0]);
            var ciphertext = Convert.FromBase64String(cookieSplit[1]);
            using (Aes aes = Aes.Create())
            {
                aes.Mode = CipherMode.CBC;
                aes.Padding = PaddingMode.PKCS7;
                aes.Key = _key;
                aes.IV = iv;

                ICryptoTransform decryptor = aes.CreateDecryptor(aes.Key, aes.IV);
                decrypted = decryptor.TransformFinalBlock(ciphertext, 0, ciphertext.Length);
            }

            return decrypted;
        }

        private static string EncryptSession(byte[] sessionData)
        {
            byte[] encrypted = null;
            byte[] iv = new byte[16];
            _rng.GetBytes(iv);

            using (Aes aes = Aes.Create())
            {
                aes.Mode = CipherMode.CBC;
                aes.Padding = PaddingMode.PKCS7;
                aes.Key = _key;
                aes.IV = iv;

                ICryptoTransform encryptor = aes.CreateEncryptor(aes.Key, aes.IV);
                encrypted = encryptor.TransformFinalBlock(sessionData, 0, sessionData.Length);
            }

            return Convert.ToBase64String(iv) + "|" + Convert.ToBase64String(encrypted);
        }

        private static byte[] SerializeSession(SecureSession session)
        {
            return Encoding.UTF8.GetBytes(JsonSerializer.Serialize(session));
        }

        private static SecureSession DeserializeSession(byte[] data)
        {
            return JsonSerializer.Deserialize<SecureSession>(data);
        }
    }
}