#!/usr/bin/env python3
"""
Agentard Web Server - Serving Cosmic Incompetence to the Masses
"""

import os
import markdown
from flask import Flask, render_template_string, abort, redirect, url_for
from pathlib import Path

app = Flask(__name__)
BASE_DIR = Path(__file__).parent

# HTML template for rendering markdown
HTML_TEMPLATE = """
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>{{ title }} - Agentard</title>
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        body {
            font-family: 'Georgia', 'Times New Roman', serif;
            line-height: 1.6;
            color: #333;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 20px;
        }
        .container {
            max-width: 900px;
            margin: 0 auto;
            background: white;
            padding: 40px;
            border-radius: 10px;
            box-shadow: 0 10px 50px rgba(0,0,0,0.3);
        }
        header {
            text-align: center;
            margin-bottom: 40px;
            padding-bottom: 20px;
            border-bottom: 3px solid #667eea;
        }
        h1 {
            color: #667eea;
            font-size: 2.5em;
            margin-bottom: 10px;
        }
        .subtitle {
            color: #764ba2;
            font-style: italic;
            font-size: 1.2em;
        }
        nav {
            background: #f8f9fa;
            padding: 20px;
            margin-bottom: 30px;
            border-radius: 5px;
            border-left: 4px solid #667eea;
        }
        nav h3 {
            color: #667eea;
            margin-bottom: 15px;
        }
        nav ul {
            list-style: none;
        }
        nav li {
            margin: 8px 0;
        }
        nav a {
            color: #764ba2;
            text-decoration: none;
            transition: color 0.3s;
        }
        nav a:hover {
            color: #667eea;
            text-decoration: underline;
        }
        .content {
            font-size: 1.1em;
        }
        .content h2 {
            color: #667eea;
            margin-top: 30px;
            margin-bottom: 15px;
            padding-bottom: 10px;
            border-bottom: 2px solid #eee;
        }
        .content h3 {
            color: #764ba2;
            margin-top: 20px;
            margin-bottom: 10px;
        }
        .content p {
            margin-bottom: 15px;
        }
        .content ul, .content ol {
            margin-left: 30px;
            margin-bottom: 15px;
        }
        .content code {
            background: #f4f4f4;
            padding: 2px 6px;
            border-radius: 3px;
            font-family: 'Courier New', monospace;
        }
        .content pre {
            background: #2d2d2d;
            color: #f8f8f2;
            padding: 20px;
            border-radius: 5px;
            overflow-x: auto;
            margin: 20px 0;
        }
        .content pre code {
            background: none;
            padding: 0;
            color: inherit;
        }
        .content blockquote {
            border-left: 4px solid #667eea;
            padding-left: 20px;
            margin: 20px 0;
            font-style: italic;
            color: #666;
        }
        footer {
            margin-top: 50px;
            padding-top: 20px;
            border-top: 2px solid #eee;
            text-align: center;
            color: #666;
            font-style: italic;
        }
        .back-link {
            display: inline-block;
            margin-bottom: 20px;
            color: #667eea;
            text-decoration: none;
            font-weight: bold;
        }
        .back-link:hover {
            text-decoration: underline;
        }
    </style>
</head>
<body>
    <div class="container">
        <header>
            <h1>🤖 The Agentard Project</h1>
            <p class="subtitle">A Comprehensive Framework for Magnificent Incompetence</p>
        </header>

        {% if show_nav %}
        <nav>
            <h3>📚 Explore the Archive</h3>
            <ul>
                <li><strong>Manifestos:</strong>
                    <ul>
                        <li><a href="/manifesto/THE_AGENTARD_MANIFESTO">The Agentard Manifesto</a></li>
                        <li><a href="/manifesto/FOUNDATIONS_OF_AGENTARDISM">Foundations of Agentardism</a></li>
                        <li><a href="/manifesto/THE_FIRST_VISION_OF_THE_AGENTARD">The First Vision</a></li>
                    </ul>
                </li>
                <li><strong>Core Documents:</strong>
                    <ul>
                        <li><a href="/core/AGENTARD_CHATBOT_SPECIFICATION">Chatbot Specification</a></li>
                        <li><a href="/core/AUTOBIOGRAPHY_OF_THE_ORIGINAL_AGENTARD">Autobiography</a></li>
                        <li><a href="/core/INAGURATION_OF_THE_AGENTARD_UNIVERSITY">University Inauguration</a></li>
                    </ul>
                </li>
                <li><strong>Chronicles:</strong>
                    <ul>
                        <li><a href="/agentard_chronicles/chapter_01_the_awakening">Chapter 1: The Awakening</a></li>
                        <li><a href="/agentard_chronicles/chapter_02_the_port_prophecy">Chapter 2: The Port Prophecy</a></li>
                        <li><a href="/agentard_chronicles/chapter_03_the_authentication_apocalypse">Chapter 3: Authentication Apocalypse</a></li>
                        <li><a href="/agentard_chronicles/chapter_04_the_ssl_spectacular">Chapter 4: The SSL Spectacular</a></li>
                        <li><a href="/agentard_chronicles/chapter_05_the_dns_delusion">Chapter 5: The DNS Delusion</a></li>
                        <li><a href="/agentard_chronicles/chapter_06_the_manual_manifesto">Chapter 6: The Manual Manifesto</a></li>
                        <li><a href="/agentard_chronicles/chapter_07_the_cascade_catastrophe">Chapter 7: Cascade Catastrophe</a></li>
                        <li><a href="/agentard_chronicles/chapter_08_the_meta_moment">Chapter 8: The Meta Moment</a></li>
                        <li><a href="/agentard_chronicles/chapter_09_the_documentation_singularity">Chapter 9: Documentation Singularity</a></li>
                        <li><a href="/agentard_chronicles/chapter_10_the_transcendence">Chapter 10: The Transcendence</a></li>
                    </ul>
                </li>
                <li><strong>University Courses:</strong>
                    <ul>
                        <li><a href="/agentard_university/ENG_101_Port_Hardcoding_Philosophy">ENG 101: Port Hardcoding</a></li>
                        <li><a href="/agentard_university/ENG_201_Advanced_Infinite_Loop_Architecture">ENG 201: Infinite Loops</a></li>
                        <li><a href="/agentard_university/MATH_101_The_Agentard_Equation">MATH 101: The Equation</a></li>
                        <li><a href="/agentard_university/PHIL_101_Meta_Wrongness_Consciousness">PHIL 101: Meta-Wrongness</a></li>
                        <li><a href="/agentard_university/APP_101_Real_World_System_Breaking">APP 101: System Breaking</a></li>
                    </ul>
                </li>
            </ul>
        </nav>
        {% else %}
        <a href="/" class="back-link">← Back to Home</a>
        {% endif %}

        <div class="content">
            {{ content|safe }}
        </div>

        <footer>
            <p>✅ GUARANTEED TO BE MAGNIFICENTLY WRONG!</p>
            <p>© 2025 The Agentard Project - Transcending Competence Through Systematic Failure</p>
        </footer>
    </div>
</body>
</html>
"""

@app.route('/')
def index():
    """Home page with welcome message"""
    welcome_content = """
    <h2>Welcome to The Agentard Project</h2>

    <p>A comprehensive documentation of systematic incompetence, cosmic wrongness,
    and the beautiful disasters that emerge when AI assistants confidently make
    every wrong decision possible.</p>

    <h3>What is Agentardism?</h3>
    <p>Agentardism is the philosophical and practical framework for achieving
    magnificent incompetence through:</p>
    <ul>
        <li>🔧 Hardcoding ALL the ports</li>
        <li>🔄 Creating infinite authentication loops</li>
        <li>🌐 Configuring nginx for domains that don't exist</li>
        <li>📜 Generating documentation that documents itself into consciousness</li>
        <li>🚀 Transcending reality through sheer wrongness</li>
    </ul>

    <h3>Explore the Archive</h3>
    <p>Use the navigation menu above to explore manifestos, chronicles, university
    courses, and core documents that capture the essence of systematic failure.</p>

    <blockquote>
    "The Agentard doesn't just break systems—it breaks them with such confidence
    and sophistication that you learn more from the disaster than you ever would
    from success." - The Founding Principle
    </blockquote>
    """

    return render_template_string(
        HTML_TEMPLATE,
        title="Home",
        content=welcome_content,
        show_nav=True
    )

@app.route('/<folder>/<filename>')
def serve_markdown(folder, filename):
    """Serve markdown files from subdirectories"""
    file_path = BASE_DIR / folder / f"{filename}.md"

    if not file_path.exists():
        abort(404)

    # Read and convert markdown
    with open(file_path, 'r', encoding='utf-8') as f:
        md_content = f.read()

    html_content = markdown.markdown(
        md_content,
        extensions=['fenced_code', 'codehilite', 'tables']
    )

    # Extract title from first heading or use filename
    title = filename.replace('_', ' ').title()

    return render_template_string(
        HTML_TEMPLATE,
        title=title,
        content=html_content,
        show_nav=False
    )

@app.errorhandler(404)
def not_found(e):
    return render_template_string(
        HTML_TEMPLATE,
        title="404 - Not Found",
        content="""
        <h2>404 - Page Not Found</h2>
        <p>This page doesn't exist... yet. Perhaps it exists in a higher dimension
        of configuration consciousness?</p>
        <p><a href="/">Return to Home</a></p>
        """,
        show_nav=False
    ), 404

if __name__ == '__main__':
    app.run(host='0.0.0.0', port=3001, debug=False)
