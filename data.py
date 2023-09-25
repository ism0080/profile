import requests


def getProjects(type):
    basePath = 'https://admin.isaacmackle.com/api'
    res = requests.get(
        f"{basePath}/collections/projects/records?fields=title,description,image,project_type,source,target,collectionId,id,img_alt&filter=(project_type='{type}')")
    response = res.json()

    for project in response['items']:
        project['img'] = {
            "src": f"{basePath}/files/{project['collectionId']}/{project['id']}/{project['image']}", "alt": project["img_alt"]}

    return response['items']


base_url = "https://isaacmackle.com"

social_buttons = [
    {"url": "https://github.com/ism0080",
        "template": "icons/_github.html", "label": "GitHub"},
    {"url": "https://www.linkedin.com/in/ism0080",
        "template": "icons/_linkedin.html", "label": "LinkedIn"}
]

personal_projects = {
    "title": "Personal Projects",
    "projects": getProjects('personal')
}

completed_projects = {
    "title": "Completed Projects",
    "projects": getProjects('completed')
}

RESUME = {
    "title": 'Isaac Mackle',
    "blurb":
    'Software engineer specialising in React, .Net Core and AWS. Prepared to face any challenge thrown at me, I am a hardworking fast learner with proven experience to work will in a team environment.',
    "info": {
        "phone": '+6427 788 4011',
        "email": 'isaac.mackle@gmail.com',
        "location": 'Christchurch',
        "website": 'isaacmackle.com',
        "github": 'github.com/ism0080',
        "linkedin": 'linkedin.com/in/ism0080'
    },
    "experience": {
        "title": 'Experience',
        "jobs": [
            {
                "role": 'Software Engineer',
                "company": "Domino's Pizza Enterprise, Brisbane - (2019 - 2021)",
                "description":
                'I contributed towards many projects for different markets (AU, NZ, JP, NL, FR, BE, DE, DK) as a frontend/backend software developer.',
                "examples": [
                    {
                        "role": 'Graduate (2019 - 2021)',
                        "experiences": [
                            {
                                "project": 'New online ordering platform (website/ native app)',
                                "description": None,
                                "technology": 'React, React Native, TypeScript, Jest, Storybook, Monorepo, Redux'
                            },
                            {
                                "project": 'Pizza Prep',
                                "description": "Trained in a Domino's store to gain knowledge of the business and how IT improves the pizza making process.",
                                "technology": None
                            },
                            {
                                "project": 'Adyen payment service migration and report notifications',
                                "description": 'This included updating current Apple pay implementation',
                                "technology": '.Net core, C#, redis, swagger, AWS (Lambda, S3, SQS, Cloud Formation), NUnit, Azure DevOps pipelines'
                            },
                            {
                                "project": 'Online ordering website/ wrapper app maintenance',
                                "description":
                                'This included implementing Launch Darkly feature toggles and implementing Apple Sign In OAuth flow for wrapper app and web',
                                "technology":
                                'ASP .Net framework, C#, JavaScript, SASS, Azure (CosmosDB, App Insights, App services), TeamCity (pipelines), Octopus (deploying builds to Azure)'
                            }
                        ]
                    },
                    {
                        "role": 'Associate (2021)',
                        "experiences": [
                            {
                                "project": 'GPS Driver Tracker',
                                "description":
                                "Working with Domino's US to implement their GPS solution for DPE through the use of messaging systems, APIs, and frontend applications",
                                "technology":
                                '.Net core, C#, ActiveMQ, Docker, React, AWS (Kinesis Firehose, Cloudformation, ECS, CloudWatch), Azure Devops Pipelines, Application Insights'
                            }
                        ]
                    }
                ]
            },
            {
                "role": 'Sales Assistant',
                "company": 'Kmart Shirley, Christchurch (2014 - 2019)',
                "description": 'Retail job with daily customer interaction',
                "examples": []
            },
            {
                "role": 'Internship',
                "company": '2IQ/hospoIQ, Christchurch - (2017)',
                "description":
                '3 month internship developing ASP .NET web application in a team of three, working in a collaborative workspace (BizDojo)',
                "examples": []
            },
            {
                "role": 'Marketing',
                "company": 'Orange Marketing, Christchurch - (2015-2016)',
                "description": 'Part-time job creating social media advertisements for business Facebook pages',
                "examples": []
            }
        ]
    },
    "skills": {
        "title": 'Skills',
        'skills': [
            {
                "key": 'HTML5 | CSS',
                "value": 'Less, CSS Modules'
            },
            {
                "key": 'JavaScript',
                "value": 'ES6, React, TypeScript, GraphQL, Apollo, NextJS'
            },
            {
                "key": 'C#',
                "value": '.Net Core, API'
            },
            {
                "key": 'Testing',
                "value": 'Jest, React Testing Library, XUnit/NUnit'
            },
            {
                "key": 'UI | Design',
                "value": 'Storybook, Sketch, Zeplin, Figma'
            },
            {
                "key": 'Server-Side',
                "value": 'NodeJS, C#'
            },
            {
                "key": 'AWS',
                "value": 'S3, CloudFormation, Lambda, DynamoDB, Kinesis Firehose, SQS, SNS'
            },
            {
                "key": 'Other',
                "value": 'CI/CD Pipelines, Docker'
            },
            {
                "key": 'Version Control',
                "value": 'Git, GitHub, BitBucket, Azure DevOps'
            },
            {
                "key": 'Project Management',
                "value": 'Agile/Scrum, Azure Devops, Jira, Asana'
            }
        ],
        "other": {
            "title": 'Experience With',
            "skills": [
                {
                    "key": 'iOS',
                    "value": 'Swift, SwiftUI'
                },
                {
                    "key": 'Python',
                    "value": 'Django, Flask'
                }
            ]
        }
    },
    "academic": {
        "title": 'Academic Record',
        "records": [
            {
                "title": 'Bachelor of Information and Communication Technologies',
                "year": '2015 - 2017',
                "location": 'Ara Institute of Canterbury'
            }
        ]
    }
}
