#Configuration
```spl2
// Saved Search
savedsearch CSRF_Alerts
| search (http_method=POST OR http_method=GET) 
  AND (http_user_agent="*<script>*" OR http_user_agent="*<img*" OR http_user_agent="*<a*" OR http_user_agent="*<form*" OR http_user_agent="*<iframe*")
  AND (uri="*/archive.py*" OR uri="*filename=*")
  AND (form_data="*<script>alert(document.cookie)</script>*" 
       OR form_data="*<img src=x onerror=\"alert(document.cookie)\">*"
       OR form_data="*<a href=\"#\" onclick=\"alert(document.cookie)\">*"
       OR form_data="*<form action=\"https://attacker.com/csrf\" method=\"POST\">*"
       OR form_data="*<iframe src=\"https://vulnerable-site.com/archive.py?filename=../../../etc/passwd\"*"
       OR form_data="*filename=../../../etc/passwd*")

// Alert Triggers

// Critical Severity Trigger  
alert CSRF_Critical
| search `CSRF_Alerts`
| eval severity="critical"
| where count > 10
| sendemail to="security@example.com" subject="Critical CSRF Alert" body="Detected a high volume of potential CSRF attacks. Immediate action required."
cron_schedule = */5 * * * *  // Run every 5 minutes

// High Severity Trigger
alert CSRF_High  
| search `CSRF_Alerts`
| eval severity="high"
| where count > 5
| sendemail to="security@example.com" subject="High CSRF Alert" body="Detected multiple potential CSRF attacks. Investigation recommended."
cron_schedule = */10 * * * * // Run every 10 minutes

// Medium Severity Trigger
alert CSRF_Medium
| search `CSRF_Alerts` 
| eval severity="medium"
| where count >= 1
| sendemail to="security@example.com" subject="Medium CSRF Alert" body="Detected a potential CSRF attack. Please review."  
cron_schedule = */30 * * * * // Run every 30 minutes
```

#Notes
- The saved search looks for POST or GET requests containing suspicious content in the user agent, URI, or form data that may indicate CSRF attacks
- It searches for script tags, img tags, a tags, form tags, and iframes in the user agent
- Checks the URI for access to the vulnerable archive.py endpoint and passing filenames
- Searches form data for known CSRF payloads and attempts to access sensitive files
- Three alert triggers are set up for critical, high, and medium severity
- Critical alerts fire every 5 min if more than 10 matching events are found, indicating a high volume attack
- High alerts fire every 10 min if more than 5 events are found, suggesting multiple attacks
- Medium alerts fire every 30 min if 1 or more events are found to flag individual potential attacks
- Alerts are emailed to the security team with appropriate subject and body for each severity level
- The cron schedule is adjusted based on severity, with more frequent checks for higher severity

