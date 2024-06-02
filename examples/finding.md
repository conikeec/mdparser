## Finding (TypedDict)

**`id:`** 50146.0 (Unique identifier for the finding)

**`title:`** `
Cross-Site Request Forgery: CSRF Protection Disabled For HTTP Endpoint in `archive.py`
` (Original title of the finding)

**`scrubbed_title:`** Cross-Site Request Forgery (Sanitized version of the title, if needed)

**`category:`** a6-security-misconfiguration (Category or classification of the finding)

**`score:`** 5.0 (Severity or risk score of the finding)

**`source:`** introduction/playground/A9/archive.py:<module> (Code source where the finding originates.)

**`calling_sink:`** introduction/playground/A9/archive.py:<module> (Code sink called by the source.)

**`calling_sink_signature:`**  (Signature of the calling sink.)

**`calling_sink_dump:`** `
from django.http import JsonResponse
from django.views.decorators.csrf import csrf_exempt

from .main import Log


@csrf_exempt
def log_function_target(request):
    L = Log(request)
    if request.method == "GET":
        L.info("GET request")
        return JsonResponse({"message":"normal get request", "method":"get"},status = 200)
    if request.method == "POST":
        username = request.POST['username']
        password = request.POST['password']
        L.info(f"POST request with username {username} and password {password}")
        if username == "admin" and password == "admin":
            return JsonResponse({"message":"Loged in successfully", "method":"post"},status = 200)
        return JsonResponse({"message":"Invalid credentials", "method":"post"},status = 401)
    if request.method == "PUT":
        L.info("PUT request")
        return JsonResponse({"message":"success", "method":"put"},status = 200)
    if request.method == "DELETE":
        if request.user.is_authenticated:
            return JsonResponse({"message":"User is authenticated", "method":"delete"},status = 200)
        L.error("DELETE request")
        return JsonResponse({"message":"permission denied", "method":"delete"},status = 200)
    if request.method == "PATCH":
        L.info("PATCH request")
        return JsonResponse({"message":"success", "method":"patch"},status = 200)
    if request.method == "UPDATE":
        return JsonResponse({"message":"success", "method":"update"},status = 200)
    return JsonResponse({"message":"method not allowed"},status = 403)


# ======================================

import datetime


# f = open('test.log', 'a') --> use this file to log
class Log:
    def __init__(self,request):
        self.request = request

    def info(self,msg):
        now = datetime.datetime.now()
        f = open('test.log', 'a')
        f.write(f"INFO:{now}:{msg}\n")
        f.close()

    def warning(self,msg):
        now = datetime.datetime.now()
        f = open('test.log', 'a')
        f.write(f"WARNING:{now}:{msg}\n")
        f.close()

    def error(self,msg):
        now = datetime.datetime.now()
        f = open('test.log', 'a')
        f.write(f"ERROR:{now}:{msg}\n")
        f.close()

` (Dump of the calling sink's state.)

**`calling_sink_location:`** introduction/playground/A9/archive.py#L1-L62 (Location of the calling sink in code.)

**`sink:`** introduction/playground/A9/archive.py:<module> (Ultimate sink where the issue ends.)

**`cwe_ids:`** CWE-352 (Common Weakness Enumeration (CWE) IDs)

