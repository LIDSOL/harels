import re

class grammatics:
    def __init__(self,r,start,ends):
        self.r=r
        self.start=start
        self.ends=ends

    def replace(self,line,replaced):
        for hit in re.finditer(self.r,line):
            s=hit.start(1)
            e=hit.end(1)
            replaced[s-1]=self.start
            replaced[e]=self.ends
            for i in range(s,e):
                replaced[i]=line[i]
        


def gmr(file):
    text=""
    
    grammars=r'\\begin\{grammar\}(.*?)\\end\{grammar\}'
    definitions=r'\\nonterminaldef\{(.*?)\}'
    terminal=grammatics(r'\\terminal\{(.*?)\}',' \'','\' ')
    nonterminal=grammatics(r'\\nonterminal\{(.*?)\}',' ',' ')
    norm=grammatics(r'\\norm\{(.*?)\}','','')
    optionalterminal=grammatics(r'\\optional\{\\terminal\{(.*?)\}\}',' ','? ')
    optionalnonterminal=grammatics(r'\\optional\{\\nonterminal\{(.*?)\}\}',' ','? ')
    notdesired=["\code","\\",""]
    rgxs=[terminal,nonterminal,norm,optionalterminal,optionalnonterminal]
    nonterminalesdef={}
    current_nonterminal_def=None

    with open(file,'r',encoding="utf-8") as fp:
        text=fp.read()
    lines=text.split('\n')

    for grammar in re.finditer(grammars, text, re.DOTALL):
        content=grammar.group(0).split('\n')
        for line_content in content:
            definition=re.search(definitions,line_content)
            if(definition):
                current_nonterminal_def=definition.group(1)
                nonterminalesdef[current_nonterminal_def]=[]
            elif current_nonterminal_def and line_content.replace("\n","")!='':
                l=['' for i in range(len(line_content))]
                
                for rgx in rgxs:
                    rgx.replace(line=line_content,replaced=l)
                s=""
                for c in l: s+=c

                for notdesire in notdesired:
                    s=s.replace(notdesire,"")
                if(s!=''):
                    nonterminalesdef[current_nonterminal_def].append(str(s))


    for key,val in nonterminalesdef.items():
        print(f"{key}")
        for i in range(len(val)):
            if(i==0):
                print(f"    :{val[i]}")
            else:
                print(f"    |{val[i]}")
        print("    ;")

    

    

def main():
    gmr("hare-specification/language/expressions.tex")

main()