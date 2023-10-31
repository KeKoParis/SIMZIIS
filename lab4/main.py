def find_g(p: int):
    for g in range(2, p):
        powers = set()
        for i in range(1, p):
            powers.add(pow(g, i) % p)
        if len(powers) == p - 1:
            return g


def generate(g: int, p: int):
    num_1 = 26
    num_2 = 245

    a = pow(g, num_1) % p
    b = pow(g, num_2) % p

    secret_1 = pow(b, num_1) % p
    secret_2 = pow(a, num_2) % p

    return secret_1, secret_2


def main():
    p = 2111
    g = find_g(p)
    print("g =", g)

    secret_Alice, secret_Bob = generate(g, p)

    print("Alice's secret:", secret_Alice, " Bob's secret:", secret_Bob)


if __name__ == '__main__':
    main()
