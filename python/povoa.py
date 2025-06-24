import time
import random
from conect_db import get_connection

nomes = ["Alice", "Bruno", "Carlos", "Diana", "Eduardo", "Fernanda", "Gustavo", "Helena", "Igor", "Juliana"]

conn = get_connection()
if conn is None:
    exit()

cur = conn.cursor()

try:
    while True:
        nome = random.choice(nomes)
        cur.execute("INSERT INTO acessos (nome) VALUES (%s)", (nome,))
        conn.commit()
        print(f"Acesso registrado: {nome}")
        time.sleep(5)
except KeyboardInterrupt:
    print("Interrompido pelo usuário.")
finally:
    cur.close()
    conn.close()
