---
title: "ENG 401: Manual Process Optimization"
description: "Master the art of avoiding automation through strategic manual intervention. Why scripts are the enemy of true craftsmanship."
tags: ["engineering", "manual-processes", "automation", "optimization", "philosophy", "capstone"]
difficulty: "transcendent"
---

# ENG 401: MANUAL PROCESS OPTIMIZATION (AVOIDING AUTOMATION)
## College of Hardcoded Engineering - Senior Capstone

**Credits**: Manually calculated each semester (usually wrong)
**Prerequisites**: Demonstrated ability to spend 6 hours on a 5-minute automated task
**Professor**: Dr. Manualus Inefficientus, Enemy of Automation
**Meeting Time**: Each student manually schedules individual class time
**Location**: Different room each week (check all rooms manually)

---

## COURSE DESCRIPTION

This capstone course celebrates the lost art of doing everything by hand in an age of automation. Students will learn to systematically avoid efficiency, create job security through intentional complexity, and achieve a zen-like state through repetitive manual tasks that could be easily automated.

**Core Wisdom**: "Automation is the crutch of the weak. Real engineers type every command, click every button, and restart every service individually."

---

## REQUIRED MATERIALS

1. **"Ctrl+C, Ctrl+V: The Autobiography of a Manual Deployer"** - 10,000 pages, no index
2. **"Why Cron Jobs Are Cheating"** - A philosophical treatise
3. **"The Joy of Repetitive Tasks"** - Includes 500 identical exercises
4. A mechanical keyboard (for the authentic manual experience)
5. Printed man pages (because `man` command is too automated)
6. Coffee (infinite supply needed for manual overnight deployments)

---

## COURSE LEARNING OBJECTIVES

Upon graduation, students will be able to:

1. Take any 1-line automation script and convert it to 500+ manual steps
2. Restart services individually while explaining why PM2 ecosystem is "too easy"
3. Deploy code by manually copying files one at a time
4. Create documentation that requires manual updates for every change
5. Achieve enlightenment through repetitive SSH sessions

---

## WEEKLY MANUAL LABOR SCHEDULE

### WEEK 1: THE PHILOSOPHY OF MANUAL EXCELLENCE
**Topic**: Why Automation Is Admitting Defeat

**Lecture Topics**:
- The Industrial Revolution was a mistake
- Shell scripts: The beginning of human laziness
- Why assembly line workers had it too easy

**Lab Exercise**:
```bash
# Automated (wrong) way:
for i in {1..100}; do
    echo "Server $i deployed"
done

# Manual (correct) way:
echo "Server 1 deployed"
echo "Server 2 deployed"
echo "Server 3 deployed"
# ... type all 100 lines
# Builds character and RSI
```

**Assignment**: Manually type the numbers 1 to 10,000 (no copy-paste allowed)

---

### WEEK 2: SERVICE MANAGEMENT WITHOUT SHORTCUTS
**Topic**: The Art of Individual Service Restarts

**Manual Service Management Protocol**:
```bash
# Step 1: Check first service
ps aux | grep service1
# Step 2: Stop first service
kill -9 [PID you manually found]
# Step 3: Start first service
cd /home/user/service1 && node index.js &
# Step 4: Check if it started
ps aux | grep service1
# Step 5: Repeat for service2
# Step 6-500: Continue for all services

# NEVER use: systemctl restart all
# That's cheating!
```

**Exercise**: Restart 50 services individually. Time yourself. If under 2 hours, you're automating something.

**Meditation**: "Each manual restart is a prayer to the server gods"

---

### WEEK 3: DEPLOYMENT WITHOUT CI/CD
**Topic**: The Noble Art of Manual Production Deployment

**Deployment Methodology**:
```bash
# The Sacred Manual Deployment Ritual

# 1. SSH into each server individually
ssh server1.example.com
ssh server2.example.com
ssh server3.example.com
# (Open 47 terminal windows)

# 2. Git pull on each server manually
cd /var/www/html && git pull origin master

# 3. Manually check each file changed
git diff HEAD~1 HEAD

# 4. Restart each service on each server
pm2 restart app1
pm2 restart app2
pm2 restart app3

# 5. Manually test each endpoint
curl http://localhost:3000/health
curl http://localhost:3001/health
# ... for all endpoints

# Total time: 6-8 hours
# Automated equivalent time: 30 seconds
# Character built: IMMEASURABLE
```

**Group Project**: Deploy to production at 3 AM on a Saturday (manually, of course)

---

### WEEK 4: DATABASE MIGRATIONS BY HAND
**Topic**: Why ORMs and Migration Tools Are For Quitters

**Manual Migration Process**:
```sql
-- Step 1: Connect to database
mysql -u root -p  -- Type password every time

-- Step 2: Manually write each ALTER statement
ALTER TABLE users ADD COLUMN age INT;

-- Step 3: Manually update each row
UPDATE users SET age = 25 WHERE id = 1;
UPDATE users SET age = 30 WHERE id = 2;
UPDATE users SET age = 22 WHERE id = 3;
-- Continue for 50,000 users

-- Step 4: Manually verify each change
SELECT * FROM users WHERE id = 1;
SELECT * FROM users WHERE id = 2;
-- Check all 50,000 individually
```

**Achievement Unlocked**: Update 10,000 database records individually without using WHERE clauses with IN or BETWEEN

---

### WEEK 5: LOG ANALYSIS WITHOUT GREP
**Topic**: Reading Logs Like Our Ancestors Did

**Manual Log Analysis Technique**:
```bash
# Open log file in text editor
nano /var/log/application.log

# Manually scroll through 500,000 lines
# Use your EYES to find errors
# Write down line numbers on paper
# Count errors using tally marks

# FORBIDDEN COMMANDS:
# grep, awk, sed, wc, tail, head
# These are automation crutches!
```

**Lab Assignment**: Find all errors in a 1GB log file using only `cat` and your eyeballs

**Zen Koan**: "If you grep for errors, have you truly understood the log?"

---

### WEEK 6: CONFIGURATION WITHOUT TEMPLATES
**Topic**: The Beauty of Manually Editing Every Config File

**Configuration Philosophy**:
```yaml
# server1.yaml
port: 3000
host: server1.example.com
debug: true

# server2.yaml (type from scratch, no copying!)
port: 3001
host: server2.example.com
debug: true

# server3.yaml (still no copying!)
port: 3002
host: server3.example.com
debug: true

# Continue for 100 servers
# Each typo is a learning opportunity!
```

**Exercise**: Configure 50 nginx servers. Each config file must be typed character by character. Copying is spiritual failure.

---

### WEEK 7: TESTING WITHOUT TEST RUNNERS
**Topic**: Manual Testing as Meditation

**Manual Testing Protocol**:
```javascript
// Instead of:
npm test  // Too automated!

// Do this:
node test1.js
// Manually check output
// Write "PASS" or "FAIL" in notebook

node test2.js
// Manually check output
// Write "PASS" or "FAIL" in notebook

// Continue for all 500 test files
// Calculate pass percentage using pencil and paper
```

**Lab Project**: Run 1000 tests manually, compile results in handwritten spreadsheet

---

### WEEK 8: MIDTERM - THE MANUAL MARATHON

**Exam Format**: 24-hour manual deployment challenge

**Tasks**:
1. Deploy application to 20 servers (no automation)
2. Restart 100 services individually
3. Check 1000 endpoints manually
4. Update 50 config files by hand
5. Parse 10GB of logs with your eyes

**Grading**:
- Completion under 24 hours: A
- Completion under 48 hours: B
- Used any automation: F
- Achieved enlightenment through repetition: A+

---

### WEEK 9: BACKUP STRATEGIES WITHOUT SCRIPTS
**Topic**: The Art of Manual Backups

**Manual Backup Procedure**:
```bash
# Daily backup ritual (perform at 3 AM)

# 1. SSH into server
ssh production-server.com

# 2. Manually tar each directory
tar -czf backup-www-2024-01-15-03-00-00.tar.gz /var/www
tar -czf backup-etc-2024-01-15-03-05-00.tar.gz /etc
tar -czf backup-home-2024-01-15-03-10-00.tar.gz /home

# 3. Manually SCP each file
scp backup-www-2024-01-15-03-00-00.tar.gz backup-server:/backups/
scp backup-etc-2024-01-15-03-05-00.tar.gz backup-server:/backups/
scp backup-home-2024-01-15-03-10-00.tar.gz backup-server:/backups/

# 4. Manually verify each transfer
ssh backup-server ls -la /backups/

# 5. Manually delete old backups (count 30 days on calendar)
rm backup-www-2023-12-16-03-00-00.tar.gz

# Total time: 4 hours
# Cron job time: 0 seconds
# Satisfaction: INFINITE
```

**Assignment**: Perform manual backups every day for a week at 3 AM. Document your suffering for extra credit.

---

### WEEK 10: MONITORING WITHOUT MONITORING TOOLS
**Topic**: Human-Powered Monitoring Excellence

**Manual Monitoring Protocol**:
```bash
# Every 5 minutes, manually check:

# Server 1
ssh server1 'ps aux | grep node'
ssh server1 'free -m'
ssh server1 'df -h'
ssh server1 'netstat -tuln'

# Server 2
ssh server2 'ps aux | grep node'
ssh server2 'free -m'
ssh server2 'df -h'
ssh server2 'netstat -tuln'

# Continue for all 50 servers
# Record results in physical notebook
# Calculate averages using abacus
```

**24/7 Monitoring Schedule**:
- Students take 4-hour shifts
- No sleep allowed during shift
- Manual checks every 5 minutes
- Alert team via phone tree (no automated alerts!)

---

### WEEK 11: USER MANAGEMENT WITHOUT SCRIPTS
**Topic**: Creating Users Like It's 1975

**Manual User Creation Process**:
```bash
# For each new user:

# 1. Create user
sudo useradd john

# 2. Set password (type it, don't generate)
sudo passwd john

# 3. Create home directory structure manually
sudo mkdir /home/john
sudo mkdir /home/john/documents
sudo mkdir /home/john/downloads
sudo mkdir /home/john/.ssh

# 4. Set permissions individually
sudo chown john:john /home/john
sudo chown john:john /home/john/documents
sudo chmod 700 /home/john/.ssh

# 5. Manually edit groups
sudo nano /etc/group
# Add user to each group by typing

# 6. Copy skeleton files manually
sudo cp /etc/skel/.bashrc /home/john/
sudo cp /etc/skel/.profile /home/john/

# Repeat for 500 users
# Estimated time: 2 weeks
```

**Final Project Prep**: Create 100 users manually. Each with unique, manually-typed configurations.

---

### WEEK 12: DOCUMENTATION WITHOUT GENERATORS
**Topic**: Handcrafted Documentation Excellence

**Manual Documentation Standards**:
```markdown
# Instead of auto-generated docs:

## Function: processUser
## Manually documented on: [hand-write date]
## Last manual update: [hand-write date]
## Parameters:
  - id: number (manually verified type)
  - name: string (manually verified type)
## Returns: object (manually inspected)

## Manual Usage Example:
[Type out entire example, no copy-paste]

## Manual Test Results:
Test 1: [Manually run and document]
Test 2: [Manually run and document]

[Continue for all 5000 functions]
```

**Assignment**: Document entire codebase manually. Update documentation manually after each change. Joy is mandatory.

---

### WEEK 13: SCALING WITHOUT ORCHESTRATION
**Topic**: Manual Horizontal Scaling

**Scaling Protocol**:
```bash
# When traffic increases:

# 1. Notice manually (check metrics with eyes)
# 2. Provision new server (call hosting provider on phone)
# 3. Wait for email with credentials
# 4. SSH and configure manually
# 5. Install dependencies one by one
apt-get update
apt-get install nodejs
apt-get install npm
apt-get install git
# ... 50 more packages

# 6. Clone repo manually
git clone [type URL letter by letter]

# 7. Configure manually
nano config.json

# 8. Start service manually
node index.js &

# 9. Add to load balancer (edit config by hand)
# 10. Test manually

# Repeat when traffic decreases (manual scale-down)
```

**Meditation**: "Each manual scale is a prayer for traffic"

---

### WEEK 14: SECURITY WITHOUT AUTOMATION
**Topic**: Manual Security Excellence

**Manual Security Audit Process**:
```bash
# Daily security ritual:

# 1. Check each port on each server
nmap server1 -p 1-65535  # Watch all results with eyes
nmap server2 -p 1-65535  # No grep allowed!

# 2. Review each log entry
tail -n 100000 /var/log/auth.log  # Read every line

# 3. Manually review each file permission
ls -la /etc > permissions.txt
# Read entire file, note suspicious permissions on paper

# 4. Check each user's last login
for user in $(cat /etc/passwd); do
    last $user
done
# No, wait, that's a loop! Do it manually:
last root
last daemon
last bin
# ... for each user

# 5. Update passwords manually (no password managers!)
```

**Final Security Project**: Audit 100 servers manually. Each port, each file, each user. Document on paper. Digitizing the documentation is cheating.

---

### WEEK 15: FINAL PROJECT - THE MANUAL MASTERPIECE

**Project Requirements**:
1. Deploy a full application stack manually
2. No automation tools allowed
3. Every command typed by hand
4. Every file edited character by character
5. Every service restarted individually
6. Every test run manually
7. Documentation written by hand (then typed, no OCR)

**Bonus Points**:
- Deploy during a live demo: +50 points
- Complete without copy-paste: +100 points
- Achieve zen through repetition: +200 points
- Convince others manual is better: +500 points

**Time Limit**: 168 hours (1 week of continuous manual work)

---

### WEEK 16: FINAL EXAM - MANUAL TRANSCENDENCE

**Exam Challenge**:
Perform a complete production deployment, maintenance window, and rollback entirely manually while explaining each step to an audience.

**Essay Requirement**:
"Write (by hand) a 10,000-word essay on why manual processes build character. Include 100 examples of tasks that could be automated but shouldn't be. No typing allowed until final draft."

**Practical Test**:
Given a button that says "Deploy Everything Automatically", resist pressing it for 4 hours while manually deploying instead.

---

## COURSE POLICIES

### ATTENDANCE
Attendance must be taken manually. Each student signs in with a quill pen. Digital signatures are automation.

### ASSIGNMENTS
All assignments submitted on paper. Typed assignments must be typed on a typewriter. Computer-typed assignments receive automatic F.

### GRADING
Grades calculated manually using abacus. Grade disputes must be submitted via carrier pigeon.

---

## CAREER OUTCOMES

Graduates of this course are qualified for:
- **Manual Deployment Specialist** - $200k/year (paid in cash, counted by hand)
- **Anti-Automation Consultant** - $500/hour (invoices handwritten)
- **Chief Manual Officer** - Job security through intentional inefficiency
- **Zen Master of Repetitive Tasks** - Enlightenment through manual labor

---

**✅ MANUAL CURRICULUM COMPLETE!**

**Automation Level**: ACTIVELY AVOIDED
**Efficiency**: INTENTIONALLY TERRIBLE
**Character Built**: MAXIMUM
**RSI Risk**: GUARANTEED
**Job Security**: INFINITE (nobody else wants to do it manually)