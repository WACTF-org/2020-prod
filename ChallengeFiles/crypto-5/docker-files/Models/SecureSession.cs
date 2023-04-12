using System;

namespace app.Models
{
    [Serializable]
    public class SecureSession
    {
        public string Username { get; set; }
        public bool IsAdmin { get; set; }
    }
}