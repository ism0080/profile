import os
import shutil
from jinja2 import Environment, FileSystemLoader
from context import home_context, resume_context

dist = "./public"
environment = Environment(loader=FileSystemLoader("./templates/"))
home_template = environment.get_template("index.src.html")
resume_template = environment.get_template("resume.src.html")

home_filename = f"{dist}/index.html"
resume_filename = f"{dist}`/resume.html"


def CreatePublicDir():
    if not os.path.exists(dist):
        os.makedirs(dist)


def WriteTemplates():
    with open(home_filename, mode='w', encoding='utf-8') as results:
        results.write(home_template.render(home_context))
        print(f"... wrote {home_filename}")

    # with open(resume_filename, mode='w', encoding='utf-8') as results:
    #     results.write(resume_template.render(resume_context))
    #     print(f"... wrote {resume_filename}")

    os.system(
        f'tailwindcss -i ./index.css -o {dist}/styles.min.css --minify')
    print("... wrote stylesheet")

    shutil.copytree("./assets", f"{dist}/assets", dirs_exist_ok=True)
    print("... copied assets")


if __name__ == "__main__":
    CreatePublicDir()
    WriteTemplates()
