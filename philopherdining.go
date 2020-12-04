package main

/*
var wg sync.WaitGroup

type Chopstick struct{ sync.Mutex }

type Philo struct {
	Name                  int
	LeftStick, RightStick *Chopstick
}

func (p Philo) eat(wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.LeftStick.Lock()
		p.RightStick.Lock()

		fmt.Println(fmt.Sprintf("starting to eat %d", p.Name))
		time.Sleep(time.Second * 1)

		p.LeftStick.Unlock()
		p.RightStick.Unlock()

		fmt.Println(fmt.Sprintf("finishing eating %d", p.Name))
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
func main() {
	counter := 5

	chopsticks := make([]*Chopstick, counter)
	for i := 0; i < counter; i++ {
		chopsticks[i] = new(Chopstick)
	}

	wg.Add(counter)
	philos := make([]*Philo, counter)
	for i := 0; i < counter; i++ {
		philos[i] = &Philo{
			Name:       (i + 1),
			LeftStick:  chopsticks[i],
			RightStick: chopsticks[(i+1)%counter],
		}

		go philos[i].eat(&wg)
	}

	fmt.Println("Philosopher Dining Problem")

	wg.Wait()
}*/
