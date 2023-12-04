fn main() {
    let x = 5 + /* 90 + */ 5;
    println!("Is 'x' 10 or 100? x = {}", x);
    let squares: Vec<_> = (0..10).map(|i| i * i).collect();
    println!("{:?}", squares);

}
