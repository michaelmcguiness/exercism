type Color =  'black' | 'brown' | 'red' | 'orange' | 'yellow' | 'green' | 'blue' | 'violet' | 'grey' | 'white'

export class ResistorColor {
  private colors: Color[]

  constructor(colors: Color[]) {
    if (colors.length < 2) {
      throw new Error('At least two colors need to be present')
    }
    this.colors = colors
  }
  value = (): number => {
    const map = {
      "black": 0,
      "brown": 1,
      "red": 2,
      "orange": 3,
      "yellow": 4,
      "green": 5,
      "blue": 6,
      "violet": 7,
      "grey": 8,
      "white": 9
    }
    return map[this.colors[0]] * 10 + map[this.colors[1]]
  }
}
