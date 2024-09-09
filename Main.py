import pygame,random

#x and y values
x = 720
y = 720
px = 310
py = 620
cx1 = random.randint(0,640)
cy1 = random.randint(-1000,0)
cx2 = random.randint(0,640)
cy2 = random.randint(-1000,0)
cx3 = random.randint(0,640)
cy3 = random.randint(-1000,0)
tx1 = random.randint(0,640)
ty1 = random.randint(-1000,0)
tx2 = random.randint(0,640)
ty2 = random.randint(-1000,0)
lx = random.randint(0,640)
ly = random.randint(-1000,0)
ex1 = random.randint(0,640)
ey1 = random.randint(-1000,0)
ex2 = random.randint(0,640)
ey2 = random.randint(-1000,0)
ex3 = random.randint(0,640)
ey3 = random.randint(-1000,0)
ex4 = random.randint(0,640)
ey4 = random.randint(-1000,0)
ex5 = random.randint(0,640)
ey5 = random.randint(-1000,0)
e1bmbx = ex1 + 35
e1bmby = ey1
e3bmbx = ex3 + 35
e3bmby = ey3
e5bmbx = ex5 + 35
e5bmby = ey5
bx = random.randint(0,640)
by = random.randint(-1000,0)
prx = random.randint(0,640)
pry = random.randint(-1000,0)

#values
spvel = 10
vel = 5
screen = pygame.display.set_mode((x,y))
run = True
page = "menu"
clock = pygame.time.Clock()
keys = pygame.key.get_pressed()

pygame.init()
pygame.display.set_icon(pygame.image.load("Assets/32bit Squadron logo.png"))
pygame.display.set_caption("32-bit Squadron") 

while run:
    clock.tick(60)

    if page == "menu":
        #resets variables
        px = 310
        ex1 = random.randint(0,640)
        ey1 = random.randint(-1000,0)
        ex2 = random.randint(0,640)
        ey2 = random.randint(-1000,0)
        ex3 = random.randint(0,640)
        ey3 = random.randint(-1000,0)
        ex4 = random.randint(0,640)
        ey4 = random.randint(-1000,0)
        ex5 = random.randint(0,640)
        ey5 = random.randint(-1000,0)
        e1bmbx = ex1 + 35
        e1bmby = ey1
        e3bmbx = ex3 + 35
        e3bmby = ey3
        e5bmbx = ex5 + 35
        e5bmby = ey5

        #renders menu
        screen.fill((70,70,70))
        screen.blit(pygame.image.load("Assets/32bit Squadron logo.png"),(160,30))
        gamestart = screen.blit(pygame.image.load("Assets/Start.png"),(120,450))
        gamecredits = screen.blit(pygame.image.load("Assets/Credits.png"),(120,575))
        gamequit = screen.blit(pygame.image.load("Assets/Quit.png"),(380,575))

        #detects events
        for event in pygame.event.get():
            if event.type == pygame.MOUSEBUTTONDOWN and event.button == 1:
                if gamestart.collidepoint(event.pos):
                    page = "game"
                if gamecredits.collidepoint(event.pos):
                    page = "credits"
                if gamequit.collidepoint(event.pos):
                    run = False
        if event.type == pygame.QUIT:
            run = False

    if page == "credits":
        screen.fill((70,70,70))
        screen.blit(pygame.image.load("Assets/Credits Scene.png"),(50,100))
        screen.blit(pygame.image.load("Assets/32bit Squadron logo.png"),(300,40))
        screen.blit(pygame.image.load("Assets/Game Screenshot.png"),(50,450))
        gamemenu = screen.blit(pygame.image.load("Assets/Main menu.png"),(390,500))
        for event in pygame.event.get():
            if event.type == pygame.MOUSEBUTTONDOWN and event.button == 1:
                if gamemenu.collidepoint(event.pos):
                   page = "menu"
            if event.type == pygame.QUIT:
                run = False
    
    if page == "game":
        screen.fill((0,50,0))

        #renders sprites
        screen.blit(pygame.image.load("Assets/Lake.png"),(lx,ly))
        screen.blit(pygame.image.load("Assets/Tree.png"),(tx1,ty1))
        screen.blit(pygame.image.load("Assets/Tree.png"),(tx2,ty2))
        screen.blit(pygame.image.load("Assets/Base.png"),(bx,by))
        screen.blit(pygame.image.load("Assets/Parachute.png"),(prx,pry))
        player = screen.blit(pygame.image.load("Assets/Player.png"),(px,py))
        enemy1 = screen.blit(pygame.image.load("Assets/BF-109.png"),(ex1,ey1))
        e1bomb = screen.blit(pygame.image.load("Assets/bomb.png"),(e1bmbx,e1bmby))
        enemy2 = screen.blit(pygame.image.load("Assets/Avro Lancaster.png"),(ex2,ey2))
        enemy3 = screen.blit(pygame.image.load("Assets/Bomber.png"),(ex3,ey3))
        e3bomb = screen.blit(pygame.image.load("Assets/bomb.png"),(e3bmbx,e3bmby))
        enemy4 = screen.blit(pygame.image.load("Assets/Warthog.png"),(ex4,ey4))
        enemy5 = screen.blit(pygame.image.load("Assets/Tomcat.png"),(ex5,ey5))
        e5bomb = screen.blit(pygame.image.load("Assets/bomb.png"),(e5bmbx,e5bmby))
        screen.blit(pygame.image.load("Assets/Cloud 1.png"),(cx1,cy1))
        screen.blit(pygame.image.load("Assets/Cloud 2.png"),(cx2,cy2))
        screen.blit(pygame.image.load("Assets/Cloud 3.png"),(cx3,cy3))

        #player movement
        keys = pygame.key.get_pressed()
        if keys[pygame.K_d] and px < 640:
          px += spvel
        if keys[pygame.K_a] and px > 0:
          px -= spvel
        if keys[pygame.K_RIGHT] and px < 640:
          px += spvel
        if keys[pygame.K_LEFT] and px > 0:
          px -= spvel

        #player collision
        if pygame.Rect.colliderect(player,enemy1):
           page = "menu"
        if pygame.Rect.colliderect(player,e1bomb):
           page = "menu"
        if pygame.Rect.colliderect(player,enemy2):
           page = "menu"
        if pygame.Rect.colliderect(player,enemy3):
           page = "menu"
        if pygame.Rect.colliderect(player,e3bomb):
           page = "menu"
        if pygame.Rect.colliderect(player,enemy4):
           page = "menu"
        if pygame.Rect.colliderect(player,enemy5):
           page = "menu"
        if pygame.Rect.colliderect(player,e5bomb):
           page = "menu"

        #cloud movement
        if cy1 > 720:
            cx1 = random.randint(0,640)
            cy1 = random.randint(-1000,0)
        else:
           cy1 += vel
        if cy2 > 720:
           cx2 = random.randint(0,640)
           cy2 = random.randint(-1000,0)
        else:
           cy2 += vel
        if cy3 > 720:
           cx3 = random.randint(0,640)
           cy3 = random.randint(-1000,0)
        else:
           cy3 += vel

        #tree movement
        if ty1 > 720:
            tx1 = random.randint(0,640)
            ty1 = random.randint(-1000,0)
        else:
           ty1 += vel
        if ty2 > 720:
            tx2 = random.randint(0,640)
            ty2 = random.randint(-1000,0)
        else:
            ty2 += vel

        #Lake movement
        if ly > 720:
            lx = random.randint(0,640)
            ly = random.randint(-1000,0)
        else:
            ly += vel

        #Base movement
        if by > 720:
            bx = random.randint(0,640)
            by = random.randint(-1000,0)
        else:
            by += vel

        #parachute movement
        if pry > 720:
            prx = random.randint(0,640)
            pry = random.randint(-1000,0)
        else:
            pry += vel

        #enemy movement
        if ey1 > 720:
            ex1 = random.randint(0,640)
            ey1 = random.randint(-1000,0)
            e1bmbx = ex1 + 34
            e1bmby = ey1 + 50
        else:
            ey1 += spvel
            e1bmby += 15
        if ey2 > 720:
            ex2 = random.randint(0,640)
            ey2 = random.randint(-1000,0)
        else:
            ey2 += spvel
        if ey3 > 720:
            ex3 = random.randint(0,640)
            ey3 = random.randint(-1000,0)
            e3bmbx = ex3 + 34
            e3bmby = ey3 + 50
        else:
            ey3 += spvel
            e3bmby += 15
        if ey4 > 720:
            ex4 = random.randint(0,640)
            ey4 = random.randint(-1000,0)
        else:
           ey4 += spvel
        if ey5 > 720:
            ex5 = random.randint(0,640)
            ey5 = random.randint(-1000,0)
            e5bmbx = ex5 + 34
            e5bmby = ey5 + 50
        else:
            ey5 += spvel
            e5bmby += 15
            
        for event in pygame.event.get():
            if event.type == pygame.QUIT:
                run = False

    pygame.display.update()
    
pygame.quit()
