SOUND_MAP = {3: "Pling", 5: "Plang", 7: "Plong"}


def convert(number: int):
    r = []
    for sound in SOUND_MAP:
        if number % sound == 0:
            r.append(SOUND_MAP[sound])
    if not r:
        r.append(str(number))
    return "".join(r)
