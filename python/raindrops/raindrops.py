SOUND_MAP = {3: "Pling", 5: "Plang", 7: "Plong"}


def convert(number: int) -> str:
    r = "".join([SOUND_MAP[sound] for sound in SOUND_MAP if number % sound == 0])
    return r if r else str(number)
