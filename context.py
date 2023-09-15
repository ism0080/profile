import data

home_context = {
    "title": "Isaac Mackle • Software Engineer",
    "tech_stack": "{ .Net Core • React • AWS }",
    # "navigation": [
    #     {"name": "projects", "url": "#"},
    #     {"name": "resume", "url": "./resume.html"},
    #     {"name": "contact", "url": "#"}
    # ],
    "email": "isaac.mackle@gmail.com",
    "social_buttons": data.social_buttons,
    "completed_projects": data.completed_projects,
    "personal_projects": data.personal_projects
}

resume_context = {
    "navigation": [
        {"name": "home", "url": "./index.html"},
        {"name": "projects", "url": "#"},
        {"name": "contact", "url": "#"}
    ],
    "email": "isaac.mackle@gmail.com",
    "social_buttons": data.social_buttons,
    "RESUME": data.RESUME
}
