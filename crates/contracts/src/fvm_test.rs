use crate::generated::fuel_ee::ResultFilter;

#[test]
fn t() {
    let utxo_id = [1; 34];
    let owner = [2; 32];
    let asset_id = [3; 32];
    let events = ResultFilter{
        0: utxo_id.into(),
        1: owner.into(),
        2: asset_id.into(),
        3: Default::default(),
        4: Default::default(),
        5: Default::default(),
    };
    let events_string = events.to_string();
    assert_eq!(events_string, "");
}