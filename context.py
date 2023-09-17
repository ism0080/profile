import data

home_context = {
    "meta": {
        "title": "Isaac Mackle • Software Engineer",
        "description": "Software Engineer who is always looking to learn more. Front-end focused with experience in React, Typescript. Get to know me.",
        "url": data.base_url,
        "image": "./assets/meta_share.png"
    },
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
    "meta": {
        "title": "Resume • Isaac Mackle",
        "description": "Software Engineer who is always looking to learn more. Front-end focused with experience in React, Typescript. Get to know me.",
        "url": data.base_url,
        "image": "./assets/meta_share.png"
    },
    "navigation": [
        {"name": "home", "url": "./index.html"},
        {"name": "projects", "url": "#"},
        {"name": "contact", "url": "#"}
    ],
    "email": "isaac.mackle@gmail.com",
    "social_buttons": data.social_buttons,
    "RESUME": data.RESUME
}
