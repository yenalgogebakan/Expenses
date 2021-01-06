# BASIC EXPENSE CONTROL APPLICATION
## Control your expenses easily from anywere

This application is being developed in golang and GIN Web framework. All unit tests will also be written 
in golang. 

**GLC** will be de final production system but **SQLITE** will be used for local development and test.

CleanCode architectural code organization rules are strictly implemented by using interfaces.
Deneme
sdfdfd

**DOMAIN**

__USER__
userId _int64_
name _string_
surname _string_
email _string_
dateCreated _time.Time_
info _string_
gender _string_

__PASSWD__
userId _int64_
passwd _string_

__USERLOG__
userId _int64_
logDate _time.Time_
action _string_
params _string_

__EXPENSES__
userId _int64_
enxpenseId _int64_
date _time.Time_
expItem _string_
amaunt _float64_
info _string_

__SESSIONS__
sessionId _int64_
user _*user_
timeOpened _time.Time_
idleTime _time.Time_
timeClosed _time.Time_

