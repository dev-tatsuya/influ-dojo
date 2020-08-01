import 'package:influ_dojo/model/rank_user.dart';
import 'package:json_annotation/json_annotation.dart';

part 'ranking.g.dart';

@JsonSerializable(fieldRename: FieldRename.snake, explicitToJson: true)
class Ranking {
  List<RankUser> rankUsers;

  Ranking(this.rankUsers);

  factory Ranking.fromJson(Map<String, dynamic> json) => _$RankingFromJson(json);
  Map<String, dynamic> toJson() => _$RankingToJson(this);
}
