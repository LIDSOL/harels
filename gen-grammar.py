import re
with open("./attributes.tex", "r") as file:
    lines = file.read()
    r = re.compile(r'\\terminal{(.*?)}')
    for i in re.finditer(r, lines):
        print(f"\'{i.groups()[0]}\' | ")


